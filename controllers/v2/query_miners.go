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
	//response, err := client.DposStatus(ctx, request)
	// getConsensusStatus方法在proto中定义，实现在xchain
	response, err := client.GetConsensusStatus(ctx, request)
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
	type ValidatorsInfo struct {
		Validators		[]string	`json:"validators"`
		Miner			string		`json:"miner"`
		Curterm			int32		`json:"curterm"`
		//Contract 		string		`json:"contract"`
	}
	var validatorsInfo = ValidatorsInfo{}
	err = json.Unmarshal([]byte(response.ValidatorsInfo), &validatorsInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	var minersInfo = &pb.ConsensusStatus{
		Term: 			validatorsInfo.Curterm,
		Version:        response.Version,
		ConsensusName:  response.ConsensusName,
		StartHeight:    response.StartHeight,
		Miner: 			validatorsInfo.Miner,
		Validators: 	validatorsInfo.Validators,
	}
	//nodes := response.GetStatus().CheckResult
	//nodesBytes, _ := json.Marshal(nodes)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		//"resp": string(nodesBytes),
		"resp": minersInfo,
	})
}