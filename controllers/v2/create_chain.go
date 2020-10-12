package v2

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jason-cn-dev/xuper-sdk-go/account"
	"github.com/jason-cn-dev/xuper-sdk-go/pb"
	"github.com/jason-cn-dev/xuper-sdk-go/transfer"
	"github.com/jason-cn-dev/xuper-sdk-go/txhash"
	"github.com/jason-cn-dev/xupercc/conf"
	"github.com/jason-cn-dev/xupercc/controllers"
	log "github.com/jason-cn-dev/xupercc/utils"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

type Desc struct {
	Module string `json:"Module"`
	Method string `json:"Method"`
	Args   `json:"Args"`
}

type Args struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

// TODO
type ParamJSON struct {
	Version   string `json:"version"`
	Consensus struct {
		Miner string `json:"miner"`
	} `json:"consensus"`
	Predistribution []struct {
		Address string `json:"address"`
		Quota   string `json:"quota"`
	} `json:"predistribution"`
}
// TODO

func CreateChain(c *gin.Context) {

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


	args := Args{
		Name: req.Args["name"],
		Data: req.Args["data"],
	}

	desc := Desc{
		Module: "kernel",
		Method: "CreateBlockChain",
		Args:   args,
	}

	// TODO
	paramJson := []byte(args.Data)
	jsObj := &ParamJSON{}
	jsErr := json.Unmarshal(paramJson, jsObj)
	if jsErr != nil {
		return
	}
	utxoTx := &pb.Transaction{Version: 0}
	// 暂时只支持配置给一个初始化账户打钱
	for _, pd := range jsObj.Predistribution {
		amount := big.NewInt(0)
		amount.SetString(pd.Quota, 10) // 10进制转换大整数
		if amount.Cmp(big.NewInt(0)) < 0 {
			return
		}
		txOutput := &pb.TxOutput{}
		txOutput.ToAddr = []byte(pd.Address)
		txOutput.Amount = amount.Bytes()
		utxoTx.TxOutputs = append(utxoTx.TxOutputs, txOutput)

		// 暂时只支持配置给一个初始化账户打钱
		break
	}
	// TODO

	bytes, err := json.Marshal(&desc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数无效",
		})
		log.Printf("json Marshal fail, err: %s", err.Error())
		return
	}

	// TODO
	utxoTx.Desc = bytes
	utxoTx.Coinbase = true
	utxoTx.Txid, _ = txhash.MakeTransactionID(utxoTx)
	txid1, err := trans.PostTx(utxoTx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "创建平行链",
			"error": err.Error(),
		})
		log.Printf("transfer fail, err: %s", err.Error())
		return
	}
	log.Printf("create block chain successfully,txid: %s ",txid1)

	to := args.Name
	amount := "100"
	fee := "0"

	txid2, fee, err := trans.Transfer(to, amount, fee, string(bytes))
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, controllers.ErrorNotEnoughUtxo) {
			msg = "余额不足，该交易需要支付gas:" + fee + "，请修改转账金额，确保足够扣除该手续费"
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "创建平行链",
			"error": msg,
		})
		log.Printf("transfer fail, err: %s", err.Error())
		return
	}
	i, err := strconv.ParseInt(fee, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "费用转换失败",
		})
		log.Printf("strconv ParseInt fail, err: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "调用成功",
		"resp": controllers.Result{
			Txid:    txid2,
			GasUsed: i,
		},
	})
}
