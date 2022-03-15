package v1

import (
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/superconsensus-chain/xupercc/controllers"
	log "github.com/superconsensus-chain/xupercc/utils"
)

/*func QueryTxSDK(c *gin.Context) {

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

	trans := transfer.InitTrans(acc, req.Node, req.BcName)
	//给服务费用的地址
	trans.Cfg.ComplianceCheck.ComplianceCheckEndorseServiceAddr = acc.Address
	//服务地址
	trans.Cfg.EndorseServiceHost = req.Node

	tx, err := trans.QueryTx(req.Txid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "查询失败",
			"error": err.Error(),
		})
		log.Printf("query tx, err: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"resp": controllers.Result{
			Tx: tx.String(),
		},
	})
}*/

func QueryTx(c *gin.Context) {

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
		record(c, "查询交易失败", err.Error())
		log.Println("query tx: new xclient failed, error=", err)
		return
	}
	defer func() {
		closeErr := xclient.Close()
		log.Println("query block close xclient failed, error=", closeErr)
	}()

	tx, err := xclient.QueryTxByID(req.Txid, xuper.WithQueryBcname(req.BcName))
	if err != nil {
		record(c, "查询交易失败", err.Error())
		log.Println("query tx failed, error=", err)
		return
	}

	respTx := log.FullTx(tx)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "查询成功",
		"resp": respTx,
	})

	/*rawTxid, err := hex.DecodeString(req.Txid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Txid无效",
		})
		log.Printf("Txid invalid, err: %s", err.Error())
		return
	}

	txstatus := &pb.TxStatus{
		Bcname: req.BcName,
		Txid:   rawTxid,
	}

	conn, err := grpc.Dial(req.Node, grpc.WithInsecure(), grpc.WithMaxMsgSize(64<<20-1))
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()
	client := pb.NewXchainClient(conn)
	reply, err := client.QueryTx(ctx, txstatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "查询失败",
			"error": err.Error(),
		})
		log.Printf("query tx fail, err: %s", err.Error())
		return
	}

	if reply.Header.Error != pb.XChainErrorEnum_SUCCESS {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "查询失败",
			"error": reply.Header.Error.String(),
		})
		log.Printf("query tx fail, err: %s", reply.Header.Error.String())
		return
	}
	if reply.Tx == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败，找不到该交易",
		})
		log.Printf("tx not found")
		return
	}
	tx := log.FullTx(reply.Tx)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"resp": tx,
	})*/
}
