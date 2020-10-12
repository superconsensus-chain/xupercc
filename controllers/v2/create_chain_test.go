package v2

import (
	"encoding/json"
	"fmt"
	"github.com/jason-cn-dev/xupercc/controllers"
	"testing"
)

func TestCreateChain(t *testing.T) {
	req := &controllers.Req{
		Node:     "127.0.0.1:37101",
		BcName:   "xuper",
		Mnemonic: "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
		Amount:   100,
		Args: map[string]string{"name":"HelloChain","data":"{\"version\": \"1\", \"consensus\": {\"miner\":\"dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN\", \"type\":\"single\"},\"predistribution\":[{\"address\": \"gGWkpaHd2FFPkwQivifwEMdFipetzzsVr\",\"quota\": \"10000\"}],\"maxblocksize\": \"128\",\"period\": \"3000\",\"award\": \"1\"}"},
	}
	//请求参数
	jsonByte, _ := json.Marshal(req)
	fmt.Println(string(jsonByte))
}

/**
{
    "node":"127.0.0.1:37101",
    "bc_name":"xuper",
    "mnemonic":"致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
    "amount":100,
    "args":{
        "data":"{\"version\": \"1\", \"consensus\": {\"miner\":\"dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN\", \"type\":\"single\"},\"predistribution\":[{\"address\": \"gGWkpaHd2FFPkwQivifwEMdFipetzzsVr\",\"quota\": \"10000\"}],\"maxblocksize\": \"128\",\"period\": \"3000\",\"award\": \"1\"}",
        "name":"HelloChain"
    }
}
*/