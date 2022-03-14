package v2

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/superconsensus-chain/xupercc/conf"
	"github.com/superconsensus-chain/xupercc/controllers"
	log "github.com/superconsensus-chain/xupercc/utils"
	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
	"github.com/xuperchain/xuperchain/service/pb"
	"net/http"
)

type Desc struct {
	Module string `json:"Module"`
	Method string `json:"Method"`
	Args   `json:"Args"`
}

type Args struct {
	Name       string `json:"name"`
	Data       string `json:"data"`
	Group      string `json:"group"`
	Identities string `json:"identities"`
	Admin      string `json:"admin"`
}

func CreateChain(c *gin.Context) {

	req := new(controllers.Req)
	err := c.ShouldBind(req)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 400,
	//		"msg":  "参数无效",
	//	})
	//	log.Printf("param invalid, err: %s", err.Error())
	//	return
	//}

	acc, err := account.RetrieveAccount(req.Mnemonic, conf.Req.Language)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "助记词无效",
		})
		log.Printf("mnemonic can not retrieve account, err: %s", err.Error())
		return
	}
	fmt.Printf("D__acc.Address： %s \n", acc.Address)
	//trans := transfer.InitTrans(acc, req.Node, req.BcName)

	args := Args{
		Name:       req.Args["name"],
		Data:       req.Args["data"],
		Group:      req.Args["group"],
		Identities: req.Args["identities"],
		Admin:      req.Args["admin"],
	}

	//desc := Desc{
	//	Module: "kernel",
	//	//Method: "CreateBlockChain",
	//	Method: req.Args["method"],
	//	Args:   args,
	//}

	invokeRequests := &pb.InvokeRequest{
		ModuleName:   "xkernel",
		ContractName: "$parachain",
		MethodName:   req.Args["method"],
		Args:         make(map[string][]byte),
	}
	if args.Name != "" {
		invokeRequests.Args["name"] = []byte(args.Name)
	}
	if args.Data != "" {
		invokeRequests.Args["data"] = []byte(args.Data)
	}
	if args.Group != "" {
		invokeRequests.Args["group"] = []byte(args.Group)
	}
	if args.Identities != "" {
		invokeRequests.Args["identities"] = []byte(args.Identities)
	}
	if args.Admin != "" {
		invokeRequests.Args["admin"] = []byte(args.Admin)
	}
	// -100是因为conf/sdk.yaml文件中创链配置fee为100，xuper.NewRequest()参数带xuper.WithFee会额外加上yaml文件中的100
	fee := "999900"
	var  gas int64 = 1000000
	if invokeRequests.MethodName == "createChain" {
		fee = "99999999900"
		gas = 100000000000
	}
	request, err := xuper.NewRequest(acc, "xkernel", "$parachain", req.Args["method"], invokeRequests.Args, "", "", xuper.WithBcname(req.BcName), xuper.WithFee(fee))
	if err != nil {
		record(c, "创建平行链失败", err.Error())
		log.Println("new request failed, error=", err)
		return
	}
	xclient, err := xuper.New(req.Node, xuper.WithConfigFile("./conf/sdk.yaml"))
	if err != nil {
		record(c, "创建平行链失败", err.Error())
		log.Println("new xclient failed, error=", err)
		return
	}

	//	方法一
	_, err = xclient.PreExecTx(request)
	if err != nil {
		record(c, "创建平行链失败", err.Error())
		log.Println("pre exec tx failed, error=", err)
		return
	}
	postTx, err := xclient.Do(request)
	fmt.Println("do tx", err)
	if err != nil {
		return
	}

	/*// 方法二，实际上方法一最后还是会走到sdk中去掉此方法的
	configFile, err := config.GetConfig("./conf/sdk.yaml")
	if err != nil {
		record(c, err.Error())
		log.Println("get config failed", err)
		return
	}
	prop, err := xuper.NewProposal(xclient, request, configFile)
	if err != nil {
		record(c, err.Error())
		log.Println("new proposal failed", err)
		return
	}
	err = prop.PreExecWithSelectUtxo()
	if err != nil {
		record(c, err.Error())
		log.Println("pre exec select utxo error=", err)
		return
	}
	preTx, err := prop.GenCompleteTx()

	postTx, err := xclient.PostTx(preTx)
	if err != nil {
		record(c, err.Error())
		log.Println("post tx failed, error=", err)
		return
	}*/

	//_, err = json.Marshal(&desc)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 400,
	//		"msg":  "参数无效",
	//	})
	//	log.Printf("json Marshal fail, err: %s", err.Error())
	//	return
	//}

	/*to := args.Name
	amount := "0"
	if !trans.Cfg.NoFee {
		amount = trans.Cfg.MinNewChainAmount
	}
	fee := "1000000"
	if invokeRequests.MethodName == "createChain"{
		fee = "100000000000"
	}
	fmt.Printf("D__shou xu fei %s \n",fee)
	rep := new(string)
	//txid2, err := trans.Transfer(to, amount, fee, string(bytes))
	txid2, err := trans.TransferToCreateChain(to, amount, fee,invokeRequests,rep)
	//fmt.Printf("D___rep %s \n",*rep)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, controllers.ErrorNotEnoughUtxo) {
			msg = "余额不足，该交易需要支付gas:" + fee + "，请修改转账金额，确保足够扣除该手续费"
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "创建平行链",
			"error": msg,
		})
		log.Printf("transfer fail, err: %s", err.Error())
		return
	}
	i, err := strconv.ParseInt(fee, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "费用转换失败",
		})
		log.Printf("strconv ParseInt fail, err: %s", err.Error())
		return
	}*/
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "调用成功",
		"resp": controllers.Result{
			Txid:    hex.EncodeToString(postTx.Tx.Txid),
			GasUsed: gas,
			Data:    string(postTx.ContractResponse.Body),
		},
	})
}

func record(c *gin.Context, msg, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":  400,
		"msg":   msg,
		"error": err,
	})
}
