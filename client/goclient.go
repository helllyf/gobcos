package client

import (
    "context"
    "errors"
    "math/big"
    "encoding/json"
    "fmt"
	
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
    BuildTime string          `json:"Build Time"`
    BuildType string          `json:"Build Type"`
    ChainId string            `json:"Chain Id"`
    FISCOBCOSVersion string   `json:"FISCO-BCOS Version"`
    GitBranch string          `json:"Git Branch"`
    GitCommitHash string      `json:"Git Commit Hash"`
    SupportedVersion string   `json:"Supported Version"`
}

func (c clientVersion) String() string {
    return fmt.Sprintf(`{
        "Build Time":"%s"
        "Build Type":"%s"
        "Chain Id":"%s"
        "FISCO-BCOS Version":"%s"
        "Git Branch":"%s"
        "Git Commit Hash":"%s"
        "Supported Version":"%s"
        }`, c.BuildTime, c.BuildType, c.ChainId, c.FISCOBCOSVersion, c.GitBranch, c.GitCommitHash, c.SupportedVersion)
}

// GetClientVersion returns the version of FISCO BCOS running on the nodes.
func (gc * Client) GetClientVersion(ctx context.Context)(*clientVersion, error) {
    var raw json.RawMessage
    
    err := gc.c.CallContext(ctx, &raw, "getClientVersion")
    if err != nil {
		return nil, err
	} else if len(raw) == 0 {
		return nil, NotFound
    }
    
    var cv *clientVersion
    if err := json.Unmarshal(raw, &cv); err != nil {
		return nil, err
	}
    return cv, err
}

// GetBlockNumber returns the latest block height on a given groupID.
func (gc * Client) GetBlockNumber(ctx context.Context, groupID *big.Int)(string, error) {
    var bn string
    err := gc.c.CallContext(ctx, &bn, "getBlockNumber", groupID)
    return bn, err
}
