package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/superconsensus-chain/xupercc/controllers"
	log "github.com/superconsensus-chain/xupercc/utils"
	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"net/http"
)

func AddressTransfer(c *gin.Context) {
	req := new(controllers.Req)
	err := c.ShouldBind(req)
	if err != nil {
		record(c, "参数node非空即可", err.Error())
		return
	}

	var addrTrans string

	if req.Args["type"] == "x2e" {
		// xchain address transfer to eth address
		addrTrans, _, err = account.XchainToEVMAddress(req.Args["address"])
	} else if req.Args["type"] == "e2x" {
		// eth address transfer to xchain address
		addrTrans, _, err = account.EVMToXchainAddress(req.Args["address"])
	} else {
		record(c, "type:[x2e/e2x]", "转换类型错误或参数缺失")
		return
	}

	if err != nil {
		record(c, "地址转换出错", err.Error())
		log.Println("address transfer failed, error=", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "转换成功",
		"resp": controllers.Result{
			Address: addrTrans,
		},
	})
}
