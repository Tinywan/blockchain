## 安装

### 区块链

```golang
$ go get -u github.com/Tinywan/blockchain/cmd
```

### Websocket

```golang
$ go get -u github.com/Tinywan/blockchain/websocket
```

## 通过命令行执行

```golang
// 进入目录
cd github.com/Tinywan/blockchain

// Windows环境：go run .\cmd\main.go
github.com\Tinywan\blockchain> go run cmd/main.go
Index: 0
PrevBlockHash:
CurrHash: 0d8845eb2da42f75aef4ee920f644975d73347e0331d17b37209c4f32ef4867f
Data: Genesis Block
Timestamp: 1531236284

Index: 1
PrevBlockHash: 0d8845eb2da42f75aef4ee920f644975d73347e0331d17b37209c4f32ef4867f
CurrHash: b11305449703848e79f02f0ba7f7db6bdd085a4a5ea50382ea4cca77644c376b
Data: Send 1 BTC to Jacky
Timestamp: 1531236284

Index: 2
PrevBlockHash: b11305449703848e79f02f0ba7f7db6bdd085a4a5ea50382ea4cca77644c376b
CurrHash: 751c3793ee3492f5e050c6b662f4d832bc125dde0aae813147e5459abc23f29a
Data: Send 1 Eox to Jack
Timestamp: 1531236284

```

## 通过RPC 接口访问数据

#### 1、命令启动http服务 
```
// 进入目录
cd github.com/Tinywan/blockchain

// 开启http服务监听，Windows环境：go run .\rpc\Server.go 
go run rpc/Server.go
```

#### 2、创建创始区块 

URL地址：`http://localhost:8888/blockchain/get`

```json
{
    "Blocks": [
        {
            "Index": 0,
            "Timestamp": 1531235780,
            "PrevBlockHash": "",
            "Hash": "0d8845eb2da42f75aef4ee920f644975d73347e0331d17b37209c4f32ef4867f",
            "Data": "Genesis Block"
        }
    ]
}
```

#### 3、写入区块链

URL地址：`http://localhost:8888/blockchain/write?data=Send 1 Tinywan to Tinyaiai`

```json
{
    "Blocks": [
        {
            "Index": 0,
            "Timestamp": 1531235780,
            "PrevBlockHash": "",
            "Hash": "0d8845eb2da42f75aef4ee920f644975d73347e0331d17b37209c4f32ef4867f",
            "Data": "Genesis Block"
        },
        {
            "Index": 1,
            "Timestamp": 1531235893,
            "PrevBlockHash": "0d8845eb2da42f75aef4ee920f644975d73347e0331d17b37209c4f32ef4867f",
            "Hash": "b11305449703848e79f02f0ba7f7db6bdd085a4a5ea50382ea4cca77644c376b",
            "Data": "Send 1 Tinywan to Tinyaiai"
        },
        {
            "Index": 2,
            "Timestamp": 1531235909,
            "PrevBlockHash": "b11305449703848e79f02f0ba7f7db6bdd085a4a5ea50382ea4cca77644c376b",
            "Hash": "751c3793ee3492f5e050c6b662f4d832bc125dde0aae813147e5459abc23f29a",
            "Data": "Send 1 BTC to Jacky"
        }
    ]
}
```

## 步骤
#### 创建 block
#### 创建 blackchin
#### 创建 http server 
