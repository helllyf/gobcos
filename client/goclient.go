package client

import (
    "fmt"
    "context"
    "encoding/json"
    "errors"
	
    "github.com/ethereum/go-ethereum/rpc"
)

// client defines typed wrapper for the FISCO BCOS RPC API
// refer the Ethereum
type Client struct {
    c * rpc.Client 
}

// NotFound is returned by API methods if the requested item does not exist.
var NotFound = errors.New("not found")

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	return DialContext(context.Background(), rawurl)
}

func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *Client {
	return &Client{c}
}

func (gc *Client) Close() {
	gc.c.Close()
}

// Blockchain Access

type clientVersion struct {
    BuildTime string
    BuildType string
    ChainId string
    FISCOBCOSVersion string
    GitBranch string
    GitCommitHash string
    SupportedVersion string
}

// GetClientVersion returns the version of FISCO BCOS running on the nodes.
func (gc * Client) GetClientVersion(ctx context.Context)(* clientVersion, error) {
    var cv *clientVersion
    err := gc.rpcClient.CallContext(ctx, &cv, "getClientVersion")
    if err == nil && cv == nil {
        err = NotFound
    }
    return cv, err
}

// GetBlockNumber returns the latest block height on a given groupID.
func (gc * Client) GetBlockNumber(ctx context.Context, groupID )([]byte, error) {
    var raw json.RawMessage
    err := gc.rpcClient.CallContext(ctx, &raw, "getBlockNumber", 1)
    if err != nil {
        return nil,err
    }

    b,err := json.Marshal(raw)
    return b, err
}
