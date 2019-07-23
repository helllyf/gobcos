package client

// Connect creates a client that uses the given host.
func Connect(host string)( * Client, error) {
    rpcClient, err := rpc.Dial(host)
	if err != nil {
        fmt.Println("connect failed!")
        return nil, err
    }
    ethClient := ethclient.NewClient(rpcClient)
	return &Client{rpcClient, ethClient},nil
}

// main method
func main() {
    client,err := Connect("http://localhost:8545")
    if err != nil {
        fmt.Println("Connect failed!")
        return
	}
    result,err := client.GetClientVersion(context.TODO())
    num,err := client.GetBlockNumber(context.TODO())
    fmt.Println(string(result[:]))
    fmt.Println(string(num[:]))
} 