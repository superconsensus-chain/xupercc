package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuperchain/xuper-sdk-go/account"
	"github.com/xuperchain/xuper-sdk-go/contract"

	"github.com/xuperchain/xupercc/conf"
	"github.com/xuperchain/xupercc/controllers"
	log "github.com/xuperchain/xupercc/utils"
)

func GroupNode(c *gin.Context) {

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

	req.ContractName = "group_chain"

	switch req.Method {
	case "list":
		req.MethodName = "listNode"
		req.Query = true
	case "del":
		req.MethodName = "delNode"
	case "add":
		req.MethodName = "addNode"
	}

	wasmContract := contract.InitWasmContract(acc, req.Node, req.BcName, req.ContractName, req.ContractAccount)
	//给服务费用的地址
	wasmContract.Cfg.ComplianceCheck.ComplianceCheckEndorseServiceAddr = acc.Address
	//服务地址
	wasmContract.Cfg.EndorseServiceHost = req.Node

	if req.Query {
		query(c, req, wasmContract)
	} else {
		invoke(c, req, wasmContract, "listNode")
	}
}
