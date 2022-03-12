package v1

import (
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"

	"github.com/superconsensus-chain/xupercc/conf"
	"github.com/superconsensus-chain/xupercc/controllers"
	log "github.com/superconsensus-chain/xupercc/utils"
)

func MethodAcl(c *gin.Context) {

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

	//所有地址的权限都是1
	ask := make(map[string]float64)
	for _, v := range req.Address { // todo 这里只是把每个add都设置为可调用，并没有将个别剔除，待优化
		ask[v] = 1
	}
	xclient, err := xuper.New(req.Node)
	if err != nil {
		return
	}
	newacl := &xuper.ACL{
		PM: xuper.PermissionModel{
			Rule:        1,
			AcceptValue: 1.0,
		},
		AksWeight: ask,
	}
	tx, err := xclient.SetMethodACL(acc, req.ContractName, req.MethodName, newacl)
	if err != nil {
		return
	}

	/*acl := xkernel.InitAcl(acc, req.Node, req.BcName, req.ContractAccount)
	//给服务费用的地址
	acl.Cfg.ComplianceCheck.ComplianceCheckEndorseServiceAddr = acc.Address
	//服务地址
	acl.Cfg.EndorseServiceHost = req.Node

	txid, err := acl.AclDoit(xkernel.METHOD, req.ContractName, req.MethodName, ask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "设置合约权限失败",
			"error": err.Error(),
		})
		log.Printf("set method acl fail, err: %s", err.Error())
		return
	}*/

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "设置成功",
		"resp": controllers.Result{
			Txid: hex.EncodeToString(tx.Tx.Txid),
		},
	})
}
