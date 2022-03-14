package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/superconsensus-chain/xupercc/controllers"
	log "github.com/superconsensus-chain/xupercc/utils"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
	"net/http"
)

func QueryAcl(c *gin.Context) {
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

	xclient, err := xuper.New(req.Node)
	if err != nil {
		record(c, "查询acl失败", err.Error())
		log.Println("query acl: new xclient failed, error=", err)
		return
	}

	const(
		AccountSuccess = "查询合约账户acl成功"
		AccountErr = "合约账户参数不规范或不存在"
		MethodSuccess = "查询合约方法acl成功"
		MethodErr = "参数不规范或合约名/合约方法不存在"
	)
	var (
		acl = &xuper.ACL{}
		QueryMsg = AccountSuccess // 默认先是查询合约账户
		QueryErr = AccountErr
	)

	if req.ContractAccount != "" {
		acl, err = xclient.QueryAccountACL(req.ContractAccount, xuper.WithQueryBcname(req.BcName))
	}else if req.ContractName != "" && req.MethodName != "" {
		QueryMsg = MethodSuccess
		QueryErr = MethodErr
		acl, err = xclient.QueryMethodACL(req.ContractName, req.MethodName, xuper.WithQueryBcname(req.BcName))
	}else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询acl失败，检查参数",
		})
	}

	if err != nil {
		record(c, "查询acl失败", err.Error())
		log.Println("query acl failed, error=", err)
		return
	}
	if acl.AksWeight == nil {
		record(c, "查询acl失败", QueryErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  QueryMsg,
		"resp": acl,
	})

	/*aclStatus := &pb.AclStatus{
		Bcname:       req.BcName,
		AccountName:  req.ContractAccount,
		ContractName: req.ContractName,
		MethodName:   req.MethodName,
	}

	if len(aclStatus.AccountName) == 0 && len(aclStatus.ContractName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不能为空",
		})
		log.Printf("param invalid, contractAccount or contractName is empty")
		return
	}

	var action string
	if aclStatus.AccountName != "" {
		action = "查询合约账户acl"
	}

	if aclStatus.ContractName != "" && aclStatus.MethodName != "" {
		action = "查询合约方法acl"
	}
	//fmt.Println(action)

	conn, err := grpc.Dial(req.Node, grpc.WithInsecure(), grpc.WithMaxMsgSize(64<<20-1))
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()
	client := pb.NewXchainClient(conn)

	reply, err := client.QueryACL(ctx, aclStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "查询失败",
			"error": err.Error(),
		})
		log.Printf("query acl fail, err: %s", err.Error())
		return
	}

	if reply.Header.Error != pb.XChainErrorEnum_SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "查询失败",
			"error": reply.Header.Error.String(),
		})
		log.Printf("query acl fail, err: %s", reply.Header.Error.String())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  action + "成功",
		"resp": reply.GetAcl(),
	})*/
}
