# gobcos
Golang Client For FISCO BCOS

FISCO BCOS Go语言版本的SDK，整合了 FISCO BCOS 的 RPC API 服务，目前采用了以太坊的 RPC 库作为接口基础。

# 环境准备

- [Golang](https://golang.org/), 版本需不低于`1.12.6`，本项目采用`go module`进行包管理。
- [FISCO BCOS 2.0.0](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/), **需要提前运行** FISCO BCOS 区块链平台


# 测试代码

首先需要拉取代码：

```shell
git clone https://github.com/KasperLiu/gobcos.git
```

然后切换到远程`dev`分支：

```shell
git checkout -b dev origin/dev
```

进行代码测试前，请先按照实际部署的RPC URL更改`goclient_test.go`中的默认RPC连接：
```go
func GetClient(t *testing.T) (*Client) {
	// RPC API
	c, err := Dial("http://localhost:8545") // your RPC API
	if err != nil {
		t.Fatalf("can not dial to the RPC API: %v", err)
	}
	return c
}
```
测试代码默认测试的函数为`getClientVersion和getBlockNumber`，其余函数需去除注释并更改为实际存在的数据后才能执行。

执行测试代码的命令为：

```shell
go test -v -count=1 ./client
```

