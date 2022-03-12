package xkernel

import (
	"encoding/json"
	"fmt"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
	"testing"
)

func TestAcl(t *testing.T) {
	acl := xuper.ACL{}
	//acl := new(Acl)
	data := `
{
	"pm":{
		"rule":1,
		"acceptValue":1.0
	},
	"aksWeight":{
		"ZUjrEbucZYBxF6U7YJKCuSJYbBQewAMWt":1.0
	}
}
`
	err := json.Unmarshal([]byte(data), acl)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", acl)

	jsonByte, err := json.Marshal(acl)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(jsonByte))
}
