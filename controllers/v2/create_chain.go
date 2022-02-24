package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuperchain/xuper-sdk-go/account"
	"github.com/xuperchain/xuper-sdk-go/pb"
	"github.com/xuperchain/xuper-sdk-go/transfer"
	"github.com/xuperchain/xupercc/conf"
	"github.com/xuperchain/xupercc/controllers"
	log "github.com/xuperchain/xupercc/utils"
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
	Group string `json:"group"`
	Identities string `json:"identities"`
	Admin string `json:"admin"`
}

func CreateChain(c *gin.Context) {

	req := new(controllers.Req)
	err := c.ShouldBind(req)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 400,
	//		"msg":  "参数无效",
	//	})
	//	log.Printf("param invalid, err: %s", err.Error())
	//	return
	//}

	acc, err := account.RetrieveAccount(req.Mnemonic, conf.Req.Language)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "助记词无效",
		})
		log.Printf("mnemonic can not retrieve account, err: %s", err.Error())
		return
	}
	fmt.Printf("D__acc.Address： %s \n",acc.Address)
	trans := transfer.InitTrans(acc, req.Node, req.BcName)


	args := Args{
		Name: req.Args["name"],
		Data: req.Args["data"],
		Group: req.Args["group"],
		Identities: req.Args["identities"],
		Admin: req.Args["admin"],
	}

	//desc := Desc{
	//	Module: "kernel",
	//	//Method: "CreateBlockChain",
	//	Method: req.Args["method"],
	//	Args:   args,
	//}

	invokeRequests := &pb.InvokeRequest{
		ModuleName: "xkernel",
		ContractName: "$parachain",
		MethodName: req.Args["method"],
		Args:       make(map[string][]byte),
	}
	if args.Name != "" {
		invokeRequests.Args["name"] = []byte(args.Name)
	}
	if args.Data != ""{
		invokeRequests.Args["data"] = []byte(args.Data)
	}
	if args.Group != "" {
		invokeRequests.Args["group"] = []byte(args.Group)
	}
	if args.Identities != ""{
		invokeRequests.Args["identities"] = []byte(args.Identities)
	}
	if args.Admin != ""{
		invokeRequests.Args["admin"] = []byte(args.Admin)
	}

	//_, err = json.Marshal(&desc)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 400,
	//		"msg":  "参数无效",
	//	})
	//	log.Printf("json Marshal fail, err: %s", err.Error())
	//	return
	//}

	to := args.Name
	amount := "0"
	if !trans.Cfg.NoFee {
		amount = trans.Cfg.MinNewChainAmount
	}
	fee := "1000000"
	if invokeRequests.MethodName == "createChain"{
		fee = "100000000000"
	}
	fmt.Printf("D__shou xu fei %s \n",fee)
	rep := new(string)
	//txid2, err := trans.Transfer(to, amount, fee, string(bytes))
	txid2, err := trans.TransferToCreateChain(to, amount, fee,invokeRequests,rep)
	//fmt.Printf("D___rep %s \n",*rep)
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
			Data: *rep,
		},
	})
}
