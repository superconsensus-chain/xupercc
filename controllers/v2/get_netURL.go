package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/superconsensus-chain/xupercc/controllers"
	log "github.com/superconsensus-chain/xupercc/utils"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
	"net/http"
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

	xclient, err := xuper.New(req.Node)
	if err != nil {
		record(c, "获取netURL失败", err.Error())
		log.Println("netURL new xclient failed, error=", err)
		return
	}
	defer func() {
		closeErr := xclient.Close()
		log.Println("query block close xclient failed, error=", closeErr)
	}()

	url, err := xclient.QueryNetURL()
	if err != nil {
		record(c, "获取netURL失败", err.Error())
		log.Println("get netURL failed, error=", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "调用成功",
		"resp": controllers.Result{
			Data:    url,
		},
	})
	return

	/*conn, err := grpc.Dial(req.Node, grpc.WithInsecure(), grpc.WithMaxMsgSize(64<<20-1))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "grpc dial error: " + err.Error(),
		})
		log.Printf("grpc dial fail, err: %s", err.Error())
		return
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()

	cl := pb.NewXchainClient(conn)
	header := &pb.Header{
		Logid:    "",
		FromNode: "",
	}
	in := &pb.CommonIn{
		Header: header,
	}
	rawUrl, err := cl.GetNetURL(ctx, in, grpc.EmptyCallOption{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "get netURL fail, error: " + err.Error(),
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
	return*/
}
