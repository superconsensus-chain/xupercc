package v1

import (
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
	"net/http"

	"github.com/superconsensus-chain/xupercc/conf"
	"github.com/superconsensus-chain/xupercc/controllers"
	log "github.com/superconsensus-chain/xupercc/utils"
)

func CreateContractAccount(c *gin.Context) {

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
		record(c, "创建合约账户失败", err.Error())
		log.Println("create contract account: new xclient failed, error=", err)
		return
	}
	defer func() {
		closeErr := xclient.Close()
		if closeErr != nil {
			log.Println("create contract account: close xclient failed, error=", closeErr)
		}
	}()

	tx, err := xclient.CreateContractAccount(acc, "XC"+req.ContractAccount+"@"+req.BcName, xuper.WithBcname(req.BcName))
	if err != nil {
		record(c, "创建合约账户失败", err.Error())
		log.Println("create contract account: xclient create contract account failed, error=", err)
		return
	}
	acl := xuper.ACL{}
	err = json.Unmarshal(tx.ContractResponse.Body, &acl)
	if err != nil {
		record(c, "创建合约账户失败", err.Error())
		log.Println("create contract account: acl unmarshal failed, error=", err)
		return
	}

	/*//有时候随机数会是0开头，fmt会截断它，所以使用for来跳过不符合长度的数
	for len(req.ContractAccount) != 16 {
		//req.ContractAccount = "1234567812345678"
		req.ContractAccount = fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000000000000))
	}
	ca := xkernel.InitAcl(acc, req.Node, req.BcName, req.ContractAccount)
	//给服务费用的地址
	ca.Cfg.ComplianceCheck.ComplianceCheckEndorseServiceAddr = acc.Address
	//服务地址
	ca.Cfg.EndorseServiceHost = req.Node

	gas, acl, txid, err := ca.CreateContractAccount()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "创建失败",
			"error": err.Error(),
		})
		log.Printf("create contract account fail, err: %s", err.Error())
		return
	}
	*/
	//log.Printf("create contract account success, txid: %s", txid)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"resp": controllers.Result{
			Txid:            hex.EncodeToString(tx.Tx.Txid),
			AccountAcl:      &acl,
			GasUsed:         tx.GasUsed,
			ContractAccount: "XC" + req.ContractAccount + "@" + req.BcName,
		},
	})
}
