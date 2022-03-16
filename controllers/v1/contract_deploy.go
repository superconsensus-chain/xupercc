package v1

import (
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"

	"github.com/superconsensus-chain/xupercc/conf"
	"github.com/superconsensus-chain/xupercc/controllers"
	log "github.com/superconsensus-chain/xupercc/utils"
)

var codeMaxSize = 2 << 20 // 2 MiB

func ContractDeploy(c *gin.Context) {

	req := new(controllers.Req)
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数无效",
		})
		log.Printf("param invalid, err: %s", err.Error())
		return
	}

	_, err = account.RetrieveAccount(req.Mnemonic, conf.Req.Language)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "助记词无效",
		})
		log.Printf("mnemonic can not retrieve account, err: %s", err.Error())
		return
	}

	//代码检查
	if !check(c, req) {
		return
	}

	//保存文件
	if !save(c, req) {
		return
	}

	//编译合约
	if !build(c, req) {
		return
	}

	//部署合约
	deploy(c, req)

}

//代码检查
func check(c *gin.Context, req *controllers.Req) bool {

	//格式检查
	if req.Runtime != "c" && req.Runtime != "go" && req.Runtime != "solidity" && req.Runtime != "c++"{
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "合约格式无效，当前只支持c和go以及solidity",
		})
		log.Printf("runtime invalid, current code runtime: %s", req.Runtime)
		return false
	}

	//名称检查
	ccname := filepath.Base(req.ContractName)
	r := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{3,15}$`)
	if !r.MatchString(ccname) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "合约名无效，只能包含字母/数字/下划线，且为字母开头，长度为4-16",
		})
		log.Printf("name invalid, current contract name: %s", req.ContractName)
		return false
	}

	//大小检查
	if req.ContractCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "没有获取到合约代码",
		})
		log.Printf("not found contract code")
		return false
	}

	if len(req.ContractCode) > codeMaxSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "合约太大了，代码应当小于2M",
		})
		log.Printf("file size too big, current code size: %d", len(req.ContractCode))
		return false
	}

	return true
}

//保存文件
func save(c *gin.Context, req *controllers.Req) bool {
	//存放目录
	err := os.MkdirAll(conf.Code.CodePath, 0755)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"msg":   "系统内部错误",
			"error": err.Error(),
		})
		pwd, pwdErr := os.Getwd()
		log.Printf("mkdir fail, err: %s, now path: ---%s---, get pwd err: %s", err.Error(), pwd, pwdErr)
		return false
	}

	//保存文件
	var filename string //文件保存名称
	if req.Runtime == "go" {
		filename = req.ContractName + ".go"
	} else if req.Runtime == "c" || req.Runtime == "c++" {
		filename = req.ContractName + ".cc"
	} else if req.Runtime == "solidity" {
		filename = req.ContractName + ".sol"
	}
	codefile := filepath.Join(conf.Code.CodePath, filename) //文件存放位置
	//codefile := conf.Code.CodePath + filename //担心CodePath会有问题，所有用filepath.Join
	code := []byte(req.ContractCode)
	err = ioutil.WriteFile(codefile, code, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"msg":   "系统内部错误",
			"error": err.Error(),
		})
		log.Printf("contract save fail, err: %s", err.Error())
		return false
	}

	log.Printf("contract saved, file: %s", codefile)
	return true
}

//编译合约
func build(c *gin.Context, req *controllers.Req) bool {
	var builtFile string
	var err error
	if req.Runtime == "go" {
		filename := req.ContractName + ".go"
		err = controllers.BuildGo(filename)
		//wasmfile = conf.Code.WasmPath + req.ContractName + ".go.wasm"
		builtFile = filepath.Join(conf.Code.WasmPath, req.ContractName+".go.native")
	} else if req.Runtime == "c" || req.Runtime == "c++" {
		filename := req.ContractName + ".cc"
		err = controllers.BuildCC(filename)
		//wasmfile = conf.Code.WasmPath + req.ContractName + ".wasm"
		builtFile = filepath.Join(conf.Code.WasmPath, req.ContractName+".wasm")
	} else if req.Runtime == "solidity" {
		err = controllers.BuildSol(req.ContractName) // 专门不带.sol后缀
		// 实际上solc编译出来的文件应该有两个，一个bin后缀，一个abi后缀，且文件
		builtFile = filepath.Join(conf.Code.WasmPath, req.ContractName+".bin")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "合约编译失败，请检查是否有误",
			"error": err.Error(),
		})
		log.Printf("build %s file fail (langType: %s), err: %s", req.ContractName, req.Runtime, err.Error())
		return false
	}

	log.Printf("file builded, file: %s", builtFile)
	return true
}

//部署合约
func deploy(c *gin.Context, req *controllers.Req) {
	//获取身份
	acc, err := account.RetrieveAccount(req.Mnemonic, conf.Req.Language)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "助记词无效",
		})
		log.Printf("mnemonic can not retrieve account, err: %s", err.Error())
		return
	}

	xclient, err := xuper.New(req.Node)
	if err != nil {
		log.Println("xupercc new xchain client err", err)
		record(c, "合约部署/升级失败", err.Error())
		return
	}
	defer func() {
		closeErr := xclient.Close()
		if closeErr != nil {
			log.Println("contract deploy: close xclient failed, error=", closeErr)
		}
	}()
	setContractE := acc.SetContractAccount(req.ContractAccount)
	if setContractE != nil {
		log.Printf("set contract account failed, error=", setContractE)
		record(c, "合约部署/升级失败", setContractE.Error())
		return
	}
	tx := &xuper.Transaction{}
	if req.Runtime == "c" || req.Runtime == "c++" {
		file := filepath.Join(conf.Code.WasmPath, req.ContractName+".wasm")
		contractCode, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("get wasm contract filed, error=%s, filepath=%s", err, file)
			record(c, "合约部署/升级失败", err.Error())
			return
		}
		if req.Upgrade {
			tx, err = xclient.UpgradeWasmContract(acc, req.ContractName, contractCode, xuper.WithBcname(req.BcName))
			if err != nil {
				log.Println("upgrade wasm contract failed, error=", err)
				record(c, "合约升级失败", err.Error())
				return
			}
		} else {
			tx, err = xclient.DeployWasmContract(acc, req.ContractName, contractCode, req.Args, xuper.WithBcname(req.BcName))
			if err != nil {
				log.Println("deploy wasm contract failed, error=", err)
				record(c, "合约部署失败", err.Error())
				return
			}
		}
	} else if req.Runtime == "go" { // 还支持java native合约，代码使用maven编译，目前就先只使用go native
		file := filepath.Join(conf.Code.WasmPath, req.ContractName+".go.native")
		contractcode, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println("get go native contract code filed, error=", err)
			record(c, "合约部署/升级失败", err.Error())
			return
		}
		if req.Upgrade {
			tx, err = xclient.UpgradeNativeContract(acc, req.ContractName, contractcode, xuper.WithBcname(req.BcName))
			if err != nil {
				log.Println("upgrade native contract failed, error=", err)
				record(c, "合约升级失败", err.Error())
				return
			}
		} else {
			tx, err = xclient.DeployNativeGoContract(acc, req.ContractName, contractcode, req.Args, xuper.WithBcname(req.BcName))
			if err != nil {
				log.Println("deploy native go contract failed, error=", err)
				record(c, "合约部署失败", err.Error())
				return
			}
		}
	} else if req.Runtime == "solidity" {
		// solidity合约暂不支持升级
		if req.Upgrade {
			log.Println("solidity合约暂不支持升级")
			record(c, "合约升级失败", "solidity合约暂不支持升级")
			return
		}
		binfile := filepath.Join(conf.Code.WasmPath, req.ContractName, req.ContractName+".bin")
		bincode, err := ioutil.ReadFile(binfile)
		if err != nil {
			log.Println("get go native contract code filed, error=", err)
			// 因为接口不规范传参导致无法部署的合约编译结果就删了
			os.RemoveAll(filepath.Join(conf.Code.WasmPath, req.ContractName))
			record(c, "合约部署失败", "注意参数规范，contract_name应与源代码中声明的主合约名一致")
			return
		}
		abifile := filepath.Join(conf.Code.WasmPath, req.ContractName, req.ContractName+".abi")
		abicode, err := ioutil.ReadFile(abifile)
		if err != nil {
			log.Println("get go native contract code filed, error=", err)
			os.RemoveAll(filepath.Join(conf.Code.WasmPath, req.ContractName))
			record(c, "合约部署失败", "注意参数规范，contract_name应与源代码中声明的主合约名一致")
			return
		}
		tx, err = xclient.DeployEVMContract(acc, req.ContractName, abicode, bincode, req.Args, xuper.WithBcname(req.BcName))
		if err != nil {
			log.Println("deploy evm contract failed, error=", err)
			record(c, "合约部署失败", err.Error())
			return
		}
	} else {
		record(c, "合约部署失败", "不支持的合约语言")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "部署成功",
		"resp": controllers.Result{
			Txid: hex.EncodeToString(tx.Tx.Txid),
			GasUsed: tx.GasUsed,
		},
	})
}

// 记录gin上下文错误
func record(c *gin.Context, msg, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":  400,
		"msg":   msg,
		"error": err,
	})
}
