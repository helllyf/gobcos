# gobcos
Golang Client For FISCO BCOS

FISCO BCOS Go语言版本的SDK，整合了FISCO BCOS的RPC API 服务，目前采用了以太坊的RPC库作为接口基础。

# 环境准备

- Golang, 版本需不低于`1.12.6`
- FISCO BCOS, 需要提前运行 FISCO BCOS 区块链平台



# 测试代码

首先需要拉取代码：

```shell
git clone https://github.com/KasperLiu/gobcos.git
```

然后切换到远程`dev`分支：

```shell
git checkout -b dev origin/dev
```

执行命令以测试代码是否正常运行：

```shell
go test -v -count=1 ./client
```

