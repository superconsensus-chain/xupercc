# 超级链API文档

## 创建账户

**功能介绍**

- 接口名称

  CreateAccount


- 功能描述

  创建一个钱包账户。

**URI**

- URI格式

  POST /v1/create_account

**请求消息**

- 参数说明

  | 名称       | 是否必选 | 参数类型 | 说明                 |
  | ---------- | -------- | -------- | -------------------- |
  | request_id | 否       | string   | 当前请求的唯一标识id |


- 请求样例

  ```
  POST https://{ip:port}/v1/create_account
  ```

  ```json
  json 类型的请求数据
  
  {"request_id":"uuid"}
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |
  | mnemonic   | string   | 助记词               |
  | address    | string   | 该账户的钱包地址     |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"创建成功",
      "error":"",
      "resp":{
          "mnemonic":"助记词",
          "address":"该账户的钱包地址"
      }
  }
  ```



## 创建合约账户

**功能介绍**

- 接口名称

  CreateContractAccount


- 功能描述

  创建一个合约账户。

**URI**

- URI格式

  POST /v1/create_contract_account

**请求消息**

- 参数说明

  | 名称             | 是否必选 | 参数类型 | 说明                    |
  | ---------------- | -------- | -------- | ----------------------- |
  | request_id       | 否       | string   | 当前请求的唯一标识id    |
  | mnemonic         | 是       | string   | 助记词                  |
  | node             | 是       | string   | 节点ip；格式为：ip:port |
  | bc_name          | 是       | string   | 链名称                  |
  | contract_account | 否       | string   | 16位的合约账户id        |


- 请求样例

  ```
  POST https://{ip:port}/v1/create_contract_account
  ```

  ```json
  json 类型的请求数据
  
  {
      "mnemonic":"助记词",
      "node":"节点ip",
      "bc_name":"链名称",
      "contract_account":"16位的合约账户id"
  }
  
  //实例
  //创建随机账户
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南"
  }
  
  //创建自定义账户
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "1234567891234567"
  }
  ```

**响应消息**

- 要素说明

  | 名称             | 参数类型 | 说明                  |
  | ---------------- | -------- | --------------------- |
  | request_id       | string   | 当前请求的唯一标识id  |
  | code             | int      | 处理状态码            |
  | msg              | string   | 应答消息              |
  | error            | string   | 错误描述              |
  | resp             | json     | 区块链的应答数据      |
  | account_acl      | json     | 该合约账户的acl权限表 |
  | contract_account | string   | 合约账户名            |
  | pm               | json     | 权限规则              |
  | aksWeight        | json     | 权限白名单            |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"创建成功",
      "error":"",
      "resp":{
          "account_acl":{
              "pm": {
                  "rule": 1,
                  "acceptValue": 1.0
              },
              "aksWeight": {
                  "账户地址": 1.0
              }
          },
          "contract_account":"XC16位的数字@链名称"
      }
  }
  ```



## 创建合约账户（使用 desc 格式）

**功能介绍**

- 接口名称

  CreateContractAccount


- 功能描述

  创建一个合约账户。

**URI**

- URI格式

  POST /v1/create_contract_account

**请求消息**

- 参数说明

  | 名称         | 是否必选 | 参数类型          | 说明                             |
  | ------------ | -------- | ----------------- | -------------------------------- |
  | request_id   | 否       | string            | 当前请求的唯一标识id             |
  | mnemonic     | 是       | string            | 助记词                           |
  | node         | 是       | string            | 节点ip；格式为：ip:port          |
  | bc_name      | 是       | string            | 链名称                           |
  | args         | 是       | map[string]string | 使用 desc 格式创建合约地址的参数 |
  | account_name | 是       | string            | 创建指定的合约地址               |
  | acl          | 是       | string            | 访问控制列表                     |


- 请求样例

  ```
  POST https://{ip:port}/v1/create_contract_account
  ```

  ```json
  json 类型的请求数据
  
  {
      "mnemonic":"助记词",
      "node":"节点ip",
      "bc_name":"链名称",
       "args" : {
          "account_name": "合约地址",
          "acl": "访问控制列表"
      }
  }
  
  //实例
  //创建自定义账户
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "args" : {
          "account_name": "1111111111111111",
          "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 0.6},\"aksWeight\": {\"j3sjeExHbdRC2emCi3ogtaaq4ZFHSP3Jb\": 0.3,\"j3sjeExHbdRC2emCi3ogtaaq4ZFHSP3Jr\": 0.3}}"
      }
  }
  ```

**响应消息**

- 要素说明

  | 名称             | 参数类型 | 说明                  |
  | ---------------- | -------- | --------------------- |
  | request_id       | string   | 当前请求的唯一标识id  |
  | code             | int      | 处理状态码            |
  | msg              | string   | 应答消息              |
  | error            | string   | 错误描述              |
  | resp             | json     | 区块链的应答数据      |
  | account_acl      | json     | 该合约账户的acl权限表 |
  | contract_account | string   | 合约账户名            |
  | pm               | json     | 权限规则              |
  | aksWeight        | json     | 权限白名单            |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"创建成功",
      "error":"",
      "resp":{
          "account_acl":{
              "pm": {
                  "rule": 1,
                  "acceptValue": 1.0
              },
              "aksWeight": {
                  "账户地址": 1.0
              }
          },
          "contract_account":"XC16位的数字@链名称"
      }
  }
  ```

## 转账

**功能介绍**

- 接口名称

  Transfer


- 功能描述

  转账到某地址。

**URI**

- URI格式

  POST /v1/transfer

**请求消息**

- 参数说明

  | 名称       | 是否必选 | 参数类型 | 说明                    |
  | ---------- | -------- | -------- | ----------------------- |
  | request_id | 否       | string   | 当前请求的唯一标识id    |
  | mnemonic   | 是       | string   | 转账人助记词            |
  | node       | 是       | string   | 节点ip；格式为：ip:port |
  | bc_name    | 是       | string   | 链名称                  |
  | to         | 是       | string   | 收款人地址              |
  | amount     | 是       | int      | 转账金额                |
  | fee        | 否       | int      | 手续费                  |
  | desc       | 否       | string   | 转账描述                |


- 请求样例

  ```
  POST https://{ip:port}/v1/transfer
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "mnemonic": "转账人助记词",
      "to":"收款人地址",
      "amount":"转账金额"
  }
  
  //实例
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "to": "ZUjrEbucZYBxF6U7YJKCuSJYbBQewAMWt",
      "amount": 1000000
  }
  ```

**响应消息**

- 要素说明

  | 名称            | 参数类型 | 说明                 |
  | --------------- | -------- | -------------------- |
  | request_id      | string   | 当前请求的唯一标识id |
  | code            | int      | 处理状态码           |
  | msg             | string   | 应答消息             |
  | error           | string   | 错误描述             |
  | resp            | json     | 区块链的应答数据     |
  | txid            | string   | 交易id               |
  | account_balance | json     | 账户余额             |
  | balance         | string   | 余额                 |
  | isFrozen        | bool     | 冻结的金额           |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"转账成功",
      "error":"",
      "resp":{
   		"txid": "xxxxxxxxxx",
          "account_balance": [
              {"balance":"0","isFrozen":true},
              {"balance":"0"}
          ]
      }
  }
  ```



## 查询余额

**功能介绍**

- 接口名称

  Balance


- 功能描述

  查询本账户的余额。

**URI**

- URI格式

  POST /v1/balance

**请求消息**

- 参数说明

  | 名称       | 是否必选 | 参数类型 | 说明                    |
  | ---------- | -------- | -------- | ----------------------- |
  | request_id | 否       | string   | 当前请求的唯一标识id    |
  | mnemonic   | 否       | string   | 助记词                  |
  | node       | 是       | string   | 节点ip；格式为：ip:port |
  | bc_name    | 是       | string   | 链名称                  |
  | account    | 是       | string   | 要查询的账户            |


- 请求样例

  ```
  POST https://{ip:port}/v1/balance
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "account": "要查询的账户"
  }
  
  //实例
  //合约账户
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "account":"XC1234567812345678@xuper"
  }
  
  //普通账户
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "account":"ZUjrEbucZYBxF6U7YJKCuSJYbBQewAMWt"
  }
  ```

**响应消息**

- 要素说明

  | 名称            | 参数类型 | 说明                 |
  | --------------- | -------- | -------------------- |
  | request_id      | string   | 当前请求的唯一标识id |
  | code            | int      | 处理状态码           |
  | msg             | string   | 应答消息             |
  | error           | string   | 错误描述             |
  | resp            | json     | 区块链的应答数据     |
  | txid            | string   | 交易id               |
  | account_balance | json     | 账户余额             |
  | balance         | string   | 余额                 |
  | isFrozen        | bool     | 是否冻结             |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"查询成功",
      "error":"",
      "resp":{
   		"account_balance": [
              {"balance":"0","isFrozen":true},
              {"balance":"0"}
          ]
      }
  }
  ```



## 查询交易

**功能介绍**

- 接口名称

  QueryTx


- 功能描述

  根据交易id查询该交易详情。

**URI**

- URI格式

  POST /v1/query_tx

**请求消息**

- 参数说明

  | 名称       | 是否必选 | 参数类型 | 说明                    |
  | ---------- | -------- | -------- | ----------------------- |
  | request_id | 否       | string   | 当前请求的唯一标识id    |
  | node       | 是       | string   | 节点ip；格式为：ip:port |
  | bc_name    | 是       | string   | 链名称                  |
  | txid       | 是       | string   | 交易id                  |


- 请求样例

  ```
  POST https://{ip:port}/v1/query_tx
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "txid":"交易id"
  }
  
  //实例
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "txid":"2f749eef25f6d9dabf197acc08e747793b663f36f4e67f820ab67e5fa65c9762"
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |
  | tx         | json     | 交易详情             |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"创建成功",
      "error":"",
      "resp":{
   		"tx": {}
      }
  }
  ```



## 查询区块

功能介绍

- 接口名称

  QueryBlock


- 功能描述

  根据交易id查询该区块详情。

URI

- URI格式

  POST /v1/query_block

请求消息

- 参数说明

  | 名称         | 是否必选 | 参数类型 | 说明                    |
  | ------------ | -------- | -------- | ----------------------- |
  | request_id   | 否       | string   | 当前请求的唯一标识id    |
  | node         | 是       | string   | 节点ip；格式为：ip:port |
  | bc_name      | 是       | string   | 链名称                  |
  | block_id     | 否       | string   | 区块id                  |
  | block_height | 否       | int      | 区块高度                |


- 请求样例

  ```
  POST https://{ip:port}/v1/query_block
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "block_id": "区块id",
      "block_height": 0 //区块高度
  }
  
  //实例
  //根据区块id
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "block_id": "6aa48e13afce1d3e68c4f5ac16297513cea3b4b941ceccb02ddeeb28b0a40b74"
  }
  
  //根据区块高度
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "block_height": 0
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |
  | block      | json     | 区块详情             |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"根据区块xx查询区块成功",
      "error":"",
      "resp":{
          block{}
  }
  ```



## 查询权限

**功能介绍**

- 接口名称

  QueryAcl


- 功能描述

  根据交易id查询该权限详情。

**URI**

- URI格式

  POST /v1/query_acl

**请求消息**

- 参数说明

  | 名称             | 是否必选 | 参数类型 | 说明                    |
  | ---------------- | -------- | -------- | ----------------------- |
  | request_id       | 否       | string   | 当前请求的唯一标识id    |
  | node             | 是       | string   | 节点ip；格式为：ip:port |
  | bc_name          | 是       | string   | 链名称                  |
  | contract_account | 否       | string   | 要查询的合约账户        |
  | contract_name    | 否       | string   | 要查询的合约            |
  | method_name      | 否       | string   | 要查询的合约方法        |


- 请求样例

  ```
  POST https://{ip:port}/v1/query_acl
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "contract_account": "合约账户",
      "contract_name": "合约",
      "method_name": "合约方法"
  }
  
  //实例
  //查询合约账户
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "contract_account": "XC1234567812345678@xuper"
  }
  
  //查询合约方法
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "contract_name": "group_chain",
      "method_name": "listNode"
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |
  | acl        | json     | 权限详情             |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"查询合约xxacl成功",
      "error":"",
      "resp":{
          "pm": {
              "rule": 1,
              "acceptValue": 1
          },
          "aksWeight": {
              "地址1": 1,
              "地址2": 1
          }
      }
  }
  ```



## 查询节点状态

**功能介绍**

- 接口名称

  Status


- 功能描述

  获取节点的当前状态。

**URI**

- URI格式

  POST /v1/status

**请求消息**

- 参数说明

  | 名称       | 是否必选 | 参数类型 | 说明                             |
  | ---------- | -------- | -------- | -------------------------------- |
  | request_id | 否       | string   | 当前请求的唯一标识id             |
  | node       | 是       | string   | 节点ip；格式为：ip:port          |
  | bc_name    | 否       | string   | 要查询的链名称，不传则查询所有链 |


- 请求样例

  ```
  POST https://{ip:port}/v1/status
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name":"链名称"
  }
  
  //实例
  {
      "node": "127.0.0.1:37102",
      "bc_name":"xuper"
  }
  ```

**响应消息**

- 要素说明

  | 名称         | 参数类型 | 说明                 |
  | ------------ | -------- | -------------------- |
  | request_id   | string   | 当前请求的唯一标识id |
  | code         | int      | 处理状态码           |
  | msg          | string   | 应答消息             |
  | error        | string   | 错误描述             |
  | resp         | json     | 区块链的应答数据     |
  | blockchains  | json     | 该节点的所有链       |
  | name         | string   | 链名                 |
  | ledger       | json     | 账本数据             |
  | rootBlockid  | string   | 该链的根区块id       |
  | tipBlockid   | string   | 该链的最新区块id     |
  | trunkHeight  | int      | 该链的最新区块高度   |
  | block        | json     | 区块数据             |
  | blockid      | string   | 区块id               |
  | preHash      | string   | 父区块哈希           |
  | proposer     | string   | 出块人地址           |
  | height       | int      | 区块高度             |
  | timestamp    | int      | 出块时间             |
  | transactions | json     | 区块包含的交易列表   |
  | txid         | string   | 交易id               |
  | txInputs     | json     | 交易输入             |
  | txOutputs    | json     | 交易输出             |
  | timestamp    | int      | 交易时间             |
  | peers        | []string | 当前连接的p2p节点    |


- 响应样例

  ```json
  {
      "code": 200,
      "msg": "查询成功",
      "resp": {
          "blockchains": [
              {"name": "平行链1"},
              {"name": "平行链2"},//其他字段如下
              {
                  "name": "xuper", //主链
                  "ledger": {
                      "rootBlockid": "根区块id",
                      "tipBlockid": "最新的区块id",
                      "trunkHeight": 422394 //最新的区块高度
                  },
                  "block": {
                      "blockid": "区块id",
                      "preHash": "父区块哈希",
                      "proposer": "出块人地址",
                      "height": 120700, //区块高度
                      "timestamp": 1589422130000327763, //出块时间
                      //区块包含的交易列表
                      "transactions": [
                          {
                              "txid": "交易id",
                              "blockid": "交易所在区块id",
                              "txInputs": null, //交易输入
                              "txOutputs": [    //交易输出
                                  {
                                      "amount": "1000000",
                                      "toAddr": "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN"
                                  }
                              ],
                              "timestamp": 1589422130000610820 //交易时间
                          }
                      ],
                      "txCount": 1 //该块包含的总交易数
                  }
              }
          ],
          //该节点正在连接的其他节点地址列表
          "peers": [
              "127.0.0.1:37102"
          ]
      }
  }
  ```



## 部署合约

**功能介绍**

- 接口名称

  ContractDeploy


- 功能描述

  部署一份合约到链上。

  通过`upgrade`字段判断当前操作是升级合约或部署合约。

**URI**

- URI格式

  POST /v1/contract_deploy

**请求消息**

- 参数说明

  | 名称                     | 是否必选 | 参数类型          | 说明                                     |
  | ------------------------ | -------- | ----------------- | ---------------------------------------- |
  | request_id               | 否       | string            | 当前请求的唯一标识id                     |
  | node                     | 是       | string            | 节点ip；格式为：ip:port                  |
  | bc_name                  | 是       | string            | 链名称                                   |
  | contract_name            | 是       | string            | 合约名称                                 |
  | contract_account         | 是       | string            | 合约账户                                 |
  | contract_code            | 是       | string            | 合约代码                                 |
  | contract_file            | 否       | file_stream       | 合约文件                                 |
  | runtime                  | 是       | string            | 合约类型；可选：c、go                    |
  | args                     | 是       | map[string]string | 合约初始化时传入的参数                   |
  | mnemonic                 | 是       | string            | 助记词                                   |
  | endorse_service_host     | 否       | string            | 开启了背书服务的节点ip；跟node一致       |
  | endorse_service_fee      | 否       | int               | 背书手续费（预留字段）                   |
  | endorse_service_fee_addr | 否       | string            | 收取手续费的地址（预留字段）             |
  | endorse_service_addr     | 否       | string            | 支付手续费的地址（预留字段）             |
  | crypto                   | 否       | string            | 与背书节点通信使用的加密协议（预留字段） |
  | fee                      | 否       | int               | 手续费（预留字段）                       |
  | upgrade                  | 否       | bool              | 是否升级；true：升级；false：部署        |


- 请求样例

  ```
  POST https://{ip:port}/v1/contract_deploy
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "要部署合约的链名",
      "mnemonic": "助记词",
      "contract_name": "合约名称",
      "contract_account": "合约账号",
      "contract_code": "合约代码",
      //初始化参数
      "args": {      
          "key1": "value1",
          "key2": "value2",
      },
      "runtime": "c",    //合约的编写语言
      "upgrade": false  //是否是升级操作
  }
  
  //伪实例
  //部署
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "XC1234567812345678@xuper",
      "contract_name": "name",
      "contract_code": "code",
      "args": {
          "key": "vaule"
      },
      "runtime": "c"
  }
  
  //升级
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "XC1234567812345678@xuper",
      "contract_name": "name",
      "contract_code": "code",
      "args": {
          "key": "vaule" //升级不会执行init操作，所以初始化传递的参数没什么用
      },
      "runtime": "c",
      "upgrade": true,
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |
  | txid       | string   | 交易id               |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"部署成功",
      "error":"",
      "resp":{
   		"txid": "xxxxxxxxxx"
      }
  }
  ```



## 调用合约

**功能介绍**

- 接口名称

  ContractInvoke


- 功能描述

  调用链上的合约方法。
  
  通过`query`字段判断当前操作是查询合约或调用合约。

**URI**

- URI格式

  POST /v1/contract_invoke

**请求消息**

- 参数说明

  | 名称                     | 是否必选 | 参数类型          | 说明                                     |
  | ------------------------ | -------- | ----------------- | ---------------------------------------- |
  | request_id               | 否       | string            | 当前请求的唯一标识id                     |
  | node                     | 是       | string            | 节点ip；格式为：ip:port                  |
  | bc_name                  | 是       | string            | 链名称                                   |
  | contract_name            | 是       | string            | 合约名称                                 |
  | contract_account         | 否       | string            | 合约账户                                 |
  | method_name              | 是       | string            | 要调用的合约的方法名称                   |
  | args                     | 是       | map[string]string | 合约初始化时传入的参数                   |
  | mnemonic                 | 是       | string            | 助记词                                   |
  | endorse_service_host     | 否       | string            | 开启了背书服务的节点ip；最好是主链       |
  | endorse_service_fee      | 否       | int               | 背书手续费（预留字段）                   |
  | endorse_service_fee_addr | 否       | string            | 收取手续费的地址（预留字段）             |
  | endorse_service_addr     | 否       | string            | 支付手续费的地址（预留字段）             |
  | crypto                   | 否       | string            | 与背书节点通信使用的加密协议（预留字段） |
  | fee                      | 否       | int               | 手续费（预留字段）                       |
  | query                    | 否       | bool              | 是否调用；true：查询，false：调用        |


- 请求样例

  ```
  POST https://{ip:port}/v1/contract_invoke
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "mnemonic": "助记词",
      "contract_name":"合约名称",
      "method_name":"合约方法名称",
      "args":{
          "key":"value"
      },
      "query":true
  }
  
  //实例
  //查询
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_name": "group_chain",
      "method_name": "listNode",
      "args": {
          "bcname": "wtf"
      },
      "query": true
  }
  
  //调用
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_name": "group_chain",
      "method_name": "addNode",
      "args": {
          "address": "address",
          "bcname": "wtf",
          "ip": "ip"
      }
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |
  | txid       | string   | 交易id               |
  | data       | string   | 调用后返回的数据     |
  | gas_used   | int      | 消耗的gas            |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"调用成功",
      "error":"",
      "resp":{
   		"txid": "xxxxxxxxxx",
          "data":"社会主义好！",
          "gas_used":0
      }
  }
  ```



## 设置合约方法的权限

**功能介绍**

- 接口名称

  MethodAcl


- 功能描述

  设置合约方法的权限。


**URI**

- URI格式

  POST /v1/method_acl

**请求消息**

- 参数说明

  | 名称             | 是否必选 | 参数类型 | 说明                     |
  | ---------------- | -------- | -------- | ------------------------ |
  | request_id       | 否       | string   | 当前请求的唯一标识id     |
  | node             | 是       | string   | 节点ip；格式为：ip:port  |
  | bc_name          | 是       | string   | 链名称                   |
  | contract_name    | 是       | string   | 合约名称                 |
  | contract_account | 是       | string   | 合约账户                 |
  | method_name      | 是       | string   | 要设置的合约的方法名称   |
  | mnemonic         | 是       | string   | 助记词                   |
  | address          | 是       | []string | 该方法允许调用的地址列表 |


- 请求样例

  ```
  POST https://{ip:port}/v1/method_acl
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "mnemonic": "助记词",
      "contract_account": "合约账号",
      "contract_name":"合约名称",
      "method_name":"合约方法名称",
      "address":["账户地址1","账户地址2"]
  }
  
  //实例
  //单个地址
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "XC1234567812345678@xuper",
      "contract_name": "group_chain",
      "method_name": "listNode",
      "address": [
          "ZUjrEbucZYBxF6U7YJKCuSJYbBQewAMWt"
      ]
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |
  | txid       | string   | 交易id               |
  | method_acl | json     | 该方法的权限表       |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"设置成功",
      "error":"",
      "resp":{
   		"txid": "xxxxxxxxxx",
          "method_acl":{
              "pm": {
                  "rule": 1,
                  "acceptValue": 1
              },
              "aksWeight": {
                  "账户地址1": 1,
                  "账户地址2": 1
              }
          }
      }
  }
  ```



## 设置合约账户的权限

**功能介绍**

- 接口名称

  AccountAcl


- 功能描述

  设置合约账户的权限。


**URI**

- URI格式

  POST /v1/account_acl

**请求消息**

- 参数说明

  | 名称             | 是否必选 | 参数类型 | 说明                     |
  | ---------------- | -------- | -------- | ------------------------ |
  | request_id       | 否       | string   | 当前请求的唯一标识id     |
  | node             | 是       | string   | 节点ip；格式为：ip:port  |
  | bc_name          | 是       | string   | 链名称                   |
  | contract_account | 是       | string   | 合约账户                 |
  | mnemonic         | 是       | string   | 助记词                   |
  | address          | 是       | []string | 该方法允许调用的地址列表 |


- 请求样例

  ```
  POST https://{ip:port}/v1/account_acl
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "mnemonic": "助记词",
      "contract_account": "合约账号",
      "address":["账户地址1","账户地址2"]
  }
  
  //实例
  //单个地址
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "XC1234567812345678@xuper",
      "address": [
          "ZUjrEbucZYBxF6U7YJKCuSJYbBQewAMWt"
      ]
  }
  ```

**响应消息**

- 要素说明

  | 名称        | 参数类型 | 说明                 |
  | ----------- | -------- | -------------------- |
  | request_id  | string   | 当前请求的唯一标识id |
  | code        | int      | 处理状态码           |
  | msg         | string   | 应答消息             |
  | error       | string   | 错误描述             |
  | resp        | json     | 区块链的应答数据     |
  | txid        | string   | 交易id               |
  | account_acl | json     | 该合约账户的权限表   |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"设置成功",
      "error":"",
      "resp":{
   		"txid": "xxxxxxxxxx",
          "account_acl":{
              "pm": {
                  "rule": 1,
                  "acceptValue": 1
              },
              "aksWeight": {
                  "xxxxxxxxxx": 1
              }
          }
      }
  }
  ```



## 群组合约（链）

**功能介绍**

- 接口名称

  GroupChain


- 功能描述

  群组对链的操作。

  通过`method`字段判断当前操作是增加/删除/查看。

**URI**

- URI格式

  POST /v1/group_chain

**请求消息**

- 参数说明

  | 名称                     | 是否必选 | 参数类型 | 说明                                     |
  | ------------------------ | -------- | -------- | ---------------------------------------- |
  | request_id               | 否       | string   | 当前请求的唯一标识id                     |
  | node                     | 是       | string   | 节点ip；格式为：ip:port                  |
  | bc_name                  | 是       | string   | 链名称                                   |
  | method                   | 是       | string   | 操作：选其一：list、add、del             |
  | args                     | 是       | string   | 平行链名称                               |
  | mnemonic                 | 是       | string   | 助记词                                   |
  | endorse_service_host     | 否       | string   | 开启了背书服务的节点ip；最好是主链       |
  | endorse_service_fee      | 否       | int      | 背书手续费（预留字段）                   |
  | endorse_service_fee_addr | 否       | string   | 收取手续费的地址（预留字段）             |
  | endorse_service_addr     | 否       | string   | 支付手续费的地址（预留字段）             |
  | crypto                   | 否       | string   | 与背书节点通信使用的加密协议（预留字段） |
  | fee                      | 否       | int      | 手续费（预留字段）                       |

  | 操作 | 说明                           |
  | ---- | ------------------------------ |
  | list | 返回群组管理中的所有平行链链名 |
  | add  | 将平行链链名添加到群组中       |
  | del  | 从群组中删除该平行链链名       |

- 请求样例

  ```
  POST https://{ip:port}/v1/group_chain
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "mnemonic": "助记词",
      "contract_account":"合约账户",
      "method":"操作名称",
      "args":{
          "bcname":"平行链名称"
      }
  }
  
  //实例
  //链列表
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "XC1234567812345678@xuper",
      "method": "list"
  }
  
  //添加链
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "XC1234567812345678@xuper",
      "args": {
          "bcname": "wtf"
      },
      "method": "add"
  }
  
  //删除链
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "XC1234567812345678@xuper",
      "args": {
          "bcname": "wtf"
      },
      "method": "del"
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |
  | txid       | string   | 交易id               |
  | data       | string   | 群组管理中的链名     |
  | gas_used   | int      | 消耗的gas            |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"调用成功",
      "error":"",
      "resp":{
   		"txid": "xxxxxxxxxx",
          "data":["平行链1","平行链2"],
          "gas_used":0
      }
  }
  ```



## 群组合约（节点）

**功能介绍**

- 接口名称

  GroupNode


- 功能描述

  群组对节点的操作。

  通过`method`字段判断当前操作是增加/删除/查看。

**URI**

- URI格式

  POST /v1/group_node

**请求消息**

- 参数说明

  | 名称                     | 是否必选 | 参数类型          | 说明                                     |
  | ------------------------ | -------- | ----------------- | ---------------------------------------- |
  | request_id               | 否       | string            | 当前请求的唯一标识id                     |
  | node                     | 是       | string            | 节点ip；格式为：ip:port                  |
  | bc_name                  | 是       | string            | 链名称                                   |
  | method                   | 是       | string            | 操作：选其一：list、add、del             |
  | args                     | 是       | map[string]string | 平行链名，p2p地址，节点账户              |
  | mnemonic                 | 是       | string            | 助记词                                   |
  | endorse_service_host     | 是       | string            | 开启了背书服务的节点ip；最好是主链       |
  | endorse_service_fee      | 否       | int               | 背书手续费（预留字段）                   |
  | endorse_service_fee_addr | 否       | string            | 收取手续费的地址（预留字段）             |
  | endorse_service_addr     | 否       | string            | 支付手续费的地址（预留字段）             |
  | crypto                   | 否       | string            | 与背书节点通信使用的加密协议（预留字段） |
  | fee                      | 否       | int               | 手续费（预留字段）                       |

  | 操作 | 说明                                |
  | ---- | ----------------------------------- |
  | list | 返回该平行链中允许同步数据的p2p地址 |
  | add  | 将p2p地址添加到平行链中             |
  | del  | 从平行链中删除该p2p地址             |

- 请求样例

  ```
  POST https://{ip:port}/v1/group_node
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "mnemonic": "助记词",
      "contract_account":"合约账户",
      "method":"操作名称",
      "args":{
          "bcname":"平行链名",
          "ip":"p2p地址",
          "address":"节点账户"
      }
  }
  
  //实例
  //查询节点列表
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "XC1234567812345678@xuper",
      "args": {
          "bcname": "xuper"
      },
      "method": "list"
  }
  
  //增加节点
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "XC1234567812345678@xuper",
      "args": {
          "bcname": "wtf",
          "ip":"ip1",
          "address":"add1"
      },
      "method": "add"
  }
  
  //删除节点
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "contract_account": "XC1234567812345678@xuper",
      "args": {
          "bcname": "wtf",
          "ip":"ip1"
      },
      "method": "del"
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                  |
  | ---------- | -------- | --------------------- |
  | request_id | string   | 当前请求的唯一标识id  |
  | code       | int      | 处理状态码            |
  | msg        | string   | 应答消息              |
  | error      | string   | 错误描述              |
  | resp       | json     | 区块链的应答数据      |
  | txid       | string   | 交易id                |
  | data       | string   | 该平行链允许的p2p地址 |
  | gas_used   | int      | 消耗的gas             |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"调用成功",
      "error":"",
      "resp":{
   		"txid": "xxxxxxxxxx",
          "data":["节点1","节点2"],
          "gas_used":0
      }
  }
  ```



## 查询最新区块和交易列表

**功能介绍**

- 接口名称

  QueryLists


- 功能描述

  获取某条链的最新区块和交易列表。

**URI**

- URI格式

  POST /v1/query_list

**请求消息**

- 参数说明

  | 名称       | 是否必选 | 参数类型 | 说明                             |
  | ---------- | -------- | -------- | -------------------------------- |
  | request_id | 否       | string   | 当前请求的唯一标识id             |
  | node       | 是       | string   | 节点ip；格式为：ip:port          |
  | bc_name    | 是       | string   | 要查询的链名称                   |


- 请求样例

  ```
  POST https://{ip:port}/v1/query_list
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name":"链名称"
  }
  
  //实例
  {
      "node": "127.0.0.1:37102",
      "bc_name":"xuper"
  }
  ```

**响应消息**

- 要素说明

  | 名称         | 参数类型 | 说明                 |
  | ------------ | -------- | -------------------- |
  | request_id   | string   | 当前请求的唯一标识id |
  | code         | int      | 处理状态码           |
  | msg          | string   | 应答消息             |
  | error        | string   | 错误描述             |
  | resp         | json数组  | 区块链的应答数据      |



- 响应样例

数据格式：json[[区块列表],[交易列表]]

  ```json
{
    "code": 200,
    "msg": "查询成功",
    "resp": [
        [
            {
                "blockid": "0000008f39c3e026afcabed1699e8aeba1b0d336785ed5225e1dd7f80410cf3d",
                "preHash": "00000115b3a875b56a58fd4882a343a6a975c0e657a88422ff9baf74d3e02c2f",
                "proposer": "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN",
                "height": 1670,
                "timestamp": 1593417852870690126,
                "transactions": [
                    {
                        "txid": "c46ce70a00530caa0eece7b34e628d39bc33dc26d20661a888a6f9662cc80efc",
                        "blockid": "0000008f39c3e026afcabed1699e8aeba1b0d336785ed5225e1dd7f80410cf3d",
                        "txOutputs": [
                            {
                                "amount": "5000000000",
                                "toAddr": "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN"
                            }
                        ],
                        "timestamp": 1593417852871140314,
                        "coinbase": true
                    }
                ],
                "txCount": 1,
                "inTrunk": true
            },
            ...
        ],
        [
            {
                "txid": "c46ce70a00530caa0eece7b34e628d39bc33dc26d20661a888a6f9662cc80efc",
                "blockid": "0000008f39c3e026afcabed1699e8aeba1b0d336785ed5225e1dd7f80410cf3d",
                "txOutputs": [
                    {
                        "amount": "5000000000",
                        "toAddr": "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN"
                    }
                ],
                "timestamp": 1593417852871140314,
                "coinbase": true
            },
            ...
        ]
    ]
}
  ```

## 创建平行链

**功能介绍**

- 接口名称

  CreateChain


- 功能描述

  创建平行链。

**URI**

- URI格式

  POST /v1/create_chain

**请求消息**

- 参数说明

  | 名称       | 是否必选 | 参数类型 | 说明                    |
  | ---------- | -------- | -------- | ----------------------- |
  | request_id | 否       | string   | 当前请求的唯一标识id    |
  | mnemonic   | 是       | string   | 转账人助记词            |
  | node       | 是       | string   | 节点ip；格式为：ip:port |
  | bc_name    | 是       | string   | 链名称                  |
  | args       | 是       | string   | 创建平行链需要的参数    |
  | amount     | 是       | int      | 转账金额                |
  
  | args | 是   | string | 创建平行链需要的参数 |
  | ---- | ---- | ------ | -------------------- |
  | name | 是   | string | 平行链的名称         |
  | data | 是   | string | 平行链的配置信息     |
  
  | data            | 是   | string | 平行链的配置信息 |
  | --------------- | ---- | ------ | ---------------- |
  | version         | 是   | string | 版本             |
  | consensus       | 是   | string | 共识配置         |
  | predistribution | 是   | string | 预分配配置       |
  | maxblocksize    | 是   | string | 区块大小限制     |
  | period          | 是   | string | 出块周期         |
  | award           | 是   | string | 出块奖励         |
  
  | consensus | 是   | string | 共识配置       |
  | --------- | ---- | ------ | -------------- |
  | miner     | 是   | string | 指定出块的地址 |
  | type      | 是   | string | 共识算法类型   |
  
  | predistribution | 是   | string | 预分配配置 |
  | --------------- | ---- | ------ | ---------- |
  | address         | 是   | string | 地址       |
  | quota           | 是   | string | 数量       |
  
  


- 请求样例

  ```
  POST https://{ip:port}/v1/create_chain
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name": "链名称",
      "mnemonic": "转账人助记词",
      "args": {
          "name": "平行链名称",
          "data": "平行链配置参数"
      },
      "amount":"设置创建平行链时最少要转多少utxo（门槛）到同链名的address"
  }
  
  //实例
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper",
      "mnemonic": "致 端 全 刘 积 旁 扰 蔬 伪 欢 近 南",
      "args": {
          "name": "HelloChain",
          "data": '{"version": "1", "consensus": {"miner":"dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN", "type":"single"},"predistribution":[{"address": "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN","quota": "1000000000000000"}],"maxblocksize": "128","period": "3000","award": "1000000"}'
      },
      "amount": 100
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |
  | txid       | string   | 交易id               |


- 响应样例

  ```json
  {
      "request_id":"uuid",
      "code":200,
      "msg":"转账成功",
      "error":"",
      "resp":{
   	  "txid": "xxxxxxxxxx",
          "gas_used":0
      }
  }
  ```

## 获取 netURL

**功能介绍**

- 接口名称

  GetNetURL


- 功能描述

​        获取节点的 netURL 地址

**URI**

- URI格式

  POST /v1/get_netURL

**请求消息**

- 参数说明

  | 名称       | 是否必选 | 参数类型 | 说明                    |
  | ---------- | -------- | -------- | ----------------------- |
  | request_id | 否       | string   | 当前请求的唯一标识id    |
  | node       | 是       | string   | 节点ip；格式为：ip:port |


- 请求样例

  ```
  POST https://{ip:port}/v1/get_netURL
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip"
  }
  
  //实例
  {
      "node": "127.0.0.1:37102"
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |
  | data       | string   | 调用后返回的数据     |


- 响应样例

  ```json
  {
      "code": 200,
      "msg": "调用成功",
      "resp": {
          "data": "/ip4/127.0.0.1/tcp/47102/p2p/QmVxeNubpg1ZQjQT8W5yZC9fD7ZB1ViArwvyGUB53sqf8e"
      }
  }
  ```

## 获取链上的矿工地址

**功能介绍**

- 接口名称

  QueryMiners


- 功能描述

​        获取链上的矿工地址

**URI**

- URI格式

  POST /v1/query_miners

**请求消息**

- 参数说明

  | 名称       | 是否必选 | 参数类型 | 说明                    |
  | ---------- | -------- | -------- | ----------------------- |
  | request_id | 否       | string   | 当前请求的唯一标识id    |
  | node       | 是       | string   | 节点ip；格式为：ip:port |
  | bc_name    | 是       | string   | 链名称                  |


- 请求样例

  ```
  POST https://{ip:port}/v1/query_miners
  ```

  ```json
  json 类型的请求数据
  
  {
      "node": "节点ip",
      "bc_name" "节点名称"
  }
  
  //实例
  {
      "node": "127.0.0.1:37102",
      "bc_name": "xuper"
  }
  ```

**响应消息**

- 要素说明

  | 名称       | 参数类型 | 说明                 |
  | ---------- | -------- | -------------------- |
  | request_id | string   | 当前请求的唯一标识id |
  | code       | int      | 处理状态码           |
  | msg        | string   | 应答消息             |
  | error      | string   | 错误描述             |
  | resp       | json     | 区块链的应答数据     |


- 响应样例

```json
{
    "code": 200,
    "msg": "查询成功",
    "resp": "[\"kKEJmXF9R4wwiYNoBsQVMhyFAR79sWR1i\",\"Te6aDDotFfKnC3uaZkuv18fneKY8uCPs5\",\"dUxWUUx9GGm49VPUt1WwwjCK9iLNet1JB\"]"
}
```
