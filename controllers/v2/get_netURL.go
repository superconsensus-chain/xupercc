package v2

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xuperchain/xuper-sdk-go/pb"
	"github.com/xuperchain/xupercc/controllers"
	log "github.com/xuperchain/xupercc/utils"
	"google.golang.org/grpc"
	"net/http"
	"time"
)


func GetNetURL(c *gin.Context) {

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
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "grpc dial error: "+err.Error(),
		})
		log.Printf("grpc dial fail, err: %s", err.Error())
		return
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()

	cl := pb.NewXchainClient(conn)
	header := &pb.Header{
		Logid:                "",
		FromNode:             "",
	}
	in := &pb.CommonIn{
		Header:               header,
	}
	rawUrl, err := cl.GetNetURL(ctx, in, grpc.EmptyCallOption{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "get netURL fail, error: "+err.Error(),
		})
		log.Printf("get netURL fail, err: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "调用成功",
		"resp": controllers.Result{
			Data:    rawUrl.RawUrl,
			Address: req.Account,
		},
	})
	return
}
