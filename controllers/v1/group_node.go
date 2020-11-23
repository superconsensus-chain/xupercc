package v1

import (
	"github.com/xuperchain/xupercc/conf"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuperchain/xuper-sdk-go/account"
	"github.com/xuperchain/xuper-sdk-go/contract"

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

	// 中间调用
	acc, err := account.GetAccountFromFile(conf.Permission.FilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无权限调用",
		})
		log.Printf("GetAccountFromFile, err: %s", err.Error())
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
