package v1

import (
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superconsensus-chain/xupercc/conf"
	"github.com/superconsensus-chain/xupercc/controllers"
	log "github.com/superconsensus-chain/xupercc/utils"
	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
)

func ContractInvoke(c *gin.Context) {
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

	xclient, err := xuper.New(req.Node)
	if err != nil {
		return
	}
	defer func() {
		closeErr := xclient.Close()
		if closeErr != nil {
			log.Println("contract invoke: close xclient failed, error=", closeErr)
		}
	}()

	tx := &xuper.Transaction{}
	if req.ModuleName == "wasm" {
		if req.Query {
			tx, err = xclient.QueryWasmContract(acc, req.ContractName, req.MethodName, req.Args, xuper.WithBcname(req.BcName))
		} else {
			tx, err = xclient.InvokeWasmContract(acc, req.ContractName, req.MethodName, req.Args, xuper.WithBcname(req.BcName))
		}
	} else if req.ModuleName == "native" {
		if req.Query {
			tx, err = xclient.QueryNativeContract(acc, req.ContractName, req.MethodName, req.Args, xuper.WithBcname(req.BcName))
		} else {
			tx, err = xclient.InvokeNativeContract(acc, req.ContractName, req.MethodName, req.Args, xuper.WithBcname(req.BcName))
		}
	} else if req.ModuleName == "evm" {
		if req.Query {
			tx, err = xclient.QueryEVMContract(acc, req.ContractName, req.MethodName, req.Args, xuper.WithBcname(req.BcName))
		} else {
			tx, err = xclient.InvokeEVMContract(acc, req.ContractName, req.MethodName, req.Args, xuper.WithBcname(req.BcName))
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code"	:	400,
			"msg"	:	"调用失败",
			"error"	:	"module_name参数缺失或格式错误",
			"hint"	:	"go合约请传native / c或c++合约传wasm / solidity合约传evm",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "操作失败",
			"error": err.Error(),
		})
		log.Printf("if query [%v] contract fail, err: %s", req.Query, err.Error())
		return
	}
	if !req.Query {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "调用成功",
			"resp": controllers.Result{
				Txid:    hex.EncodeToString(tx.Tx.Txid),
				Data:    string(tx.ContractResponse.Body),
				GasUsed: tx.GasUsed,
			},
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"resp": controllers.Result{
				Data:    string(tx.ContractResponse.Body),
			},
		})
	}
	/*
		wasmContract := contract.InitWasmContract(acc, req.Node, req.BcName, req.ContractName, req.ContractAccount)
		//给服务费用的地址
		wasmContract.Cfg.ComplianceCheck.ComplianceCheckEndorseServiceAddr = acc.Address
		//服务地址
		wasmContract.Cfg.EndorseServiceHost = req.Node

		module := req.ModuleName
		if !(module == "wasm" || module == "native" || module == "evm") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  400,
				"msg":   "查询失败",
				"error": `缺少参数module_name(值应为"wasm"或者"native")`,
			})
			return
		}
		if req.Query {
			query(c, req, wasmContract)
		} else {
			invoke(c, req, wasmContract)
		}*/
}

/*//查询的操作
func query(c *gin.Context, req *controllers.Req, wasmContract *contract.WasmContract) {
	resp, err := wasmContract.QueryWasmContract(req.MethodName, req.ModuleName, req.Args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "查询失败",
			"error": err.Error(),
		})
		log.Printf("query contract fail, err: %s", err.Error())
		return
	}

	//合约应答数据
	var datas []string
	for _, v := range resp.Response.Responses {
		//记录合约错误
		if v.Status != 200 {
			datas = append(datas, v.String())
			continue
		}

		//记录数据
		datas = append(datas, string(v.Body))
	}

	//fmt.Println("resp: ", datas)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"resp": controllers.Result{
			Data:    fmt.Sprint(datas),
			//GasUsed: resp.Response.GasUsed,
		},
	})
}

//封装的合约调用方法，sdk自带的不能获取gas
func InvokeWasmContract(c *contract.WasmContract, req *controllers.Req) (*pb.InvokeRPCResponse, string, error) {
	// preExe
	resp, err := c.PreInvokeWasmContract(req.MethodName, req.ModuleName, req.Args)
	fmt.Println("invoke args", req.Args)
	if err != nil {
		return nil, "", err
	}

	// post
	txid, err := c.PostWasmContract(resp)

	//在invoke()中获取最新的数据时，为了跟query返回的对象保持类型一致，封装成InvokeRPCResponse
	return &pb.InvokeRPCResponse{Response: resp.Response}, txid, err
}

//调用的操作
func invoke(c *gin.Context, req *controllers.Req, wasmContract *contract.WasmContract, refreshMethod ...string) {
	resp, txid, err := InvokeWasmContract(wasmContract, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "调用失败",
			"error": err.Error(),
		})
		log.Printf("invoke contract fail, err: %s", err.Error())
		return
	}

	tempResp := resp //记录旧数据
	//获取最新的数据
	if refreshMethod != nil {
		resp, err = wasmContract.QueryWasmContract(refreshMethod[0], req.ModuleName, req.Args)
		if err != nil {
			//该错误不处理，可能还没有出块，所以数据也不是最新的
			resp = tempResp
		}
	}

	//合约应答数据
	var datas []string
	for _, v := range resp.Response.Responses {
		//记录合约错误
		if v.Status != 200 {
			datas = append(datas, v.String())
			continue
		}

		//记录数据
		datas = append(datas, string(v.Body))
	}

	//fmt.Println("resp: ", datas)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "调用成功",
		"resp": controllers.Result{
			Txid:    txid,
			Data:    fmt.Sprint(datas),
			GasUsed: resp.Response.GasUsed,
		},
	})
}
*/
