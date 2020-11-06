package v2

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/xuperchain/xuper-sdk-go/pb"
	"github.com/xuperchain/xupercc/controllers"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

func QueryMiners(c *gin.Context) {

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

	conn, err := grpc.Dial(req.Node, grpc.WithInsecure(), grpc.WithMaxMsgSize(64<<20-1))
	if err != nil {
		log.Printf("can not connect to node, err: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()
	client := pb.NewXchainClient(conn)

	request := &pb.DposStatusRequest{
		Bcname: req.BcName,
	}
	response, err := client.DposStatus(ctx, request)
	if err != nil {
		log.Printf("client.DposStatus, err: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	//output, err := json.MarshalIndent(response.GetStatus(), "", "  ")
	//
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 400,
	//		"msg":  err.Error(),
	//	})
	//	return
	//}
	nodes := response.GetStatus().CheckResult
	nodesBytes, _ := json.Marshal(nodes)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"resp": string(nodesBytes),
	})
}