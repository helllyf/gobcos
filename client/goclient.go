// client provides the RPC methods for FISCO BCOS
package client

import (
    "context"
    "errors"
    "math/big"
    "encoding/json"
	
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

// Process the json result of GetClientVersion
// type clientVersion struct {
//     BuildTime string          `json:"Build Time"`
//     BuildType string          `json:"Build Type"`
//     ChainId string            `json:"Chain Id"`
//     FISCOBCOSVersion string   `json:"FISCO-BCOS Version"`
//     GitBranch string          `json:"Git Branch"`
//     GitCommitHash string      `json:"Git Commit Hash"`
//     SupportedVersion string   `json:"Supported Version"`
// }

// customized the print format of ClientVersion struct
// func (c clientVersion) String() string {
//     return fmt.Sprintf(`{
//         "Build Time":"%s"
//         "Build Type":"%s"
//         "Chain Id":"%s"
//         "FISCO-BCOS Version":"%s"
//         "Git Branch":"%s"
//         "Git Commit Hash":"%s"
//         "Supported Version":"%s"
//         }`, c.BuildTime, c.BuildType, c.ChainId, c.FISCOBCOSVersion, c.GitBranch, c.GitCommitHash, c.SupportedVersion)
// }

// GetClientVersion returns the version of FISCO BCOS running on the nodes.
func (gc *Client) GetClientVersion(ctx context.Context) ([]byte, error) {
    var raw json.RawMessage
    
    err := gc.c.CallContext(ctx, &raw, "getClientVersion")
    if err != nil {
		return nil, err
	} else if len(raw) == 0 {
		return nil, NotFound
    }
    
    // var cv *clientVersion
    // if err := json.Unmarshal(raw, &cv); err != nil {
	// 	return nil, err
	// }
    // return cv, err

    return raw,err
}

// GetBlockNumber returns the latest block height(hex format) on a given groupID.
func (gc *Client) GetBlockNumber(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getBlockNumber", groupID)
    if err != nil {
        return nil, err
    }
    return raw, err
}

// GetPBFTView returns the latest PBFT view(hex format) of the specific group and it will returns a wrong sentence
// if the consensus algorithm is not the PBFT.
func (gc *Client) GetPBFTView(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getPbftView", groupID)
    if err != nil {
        return nil, err
    }
    return raw, err

    // TODO
    // Raft consensus
}

// GetSealerList returns the list of consensus nodes' ID according to the groupID
func (gc *Client) GetSealerList(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getSealerList", groupID)
    if err != nil {
        return nil, err
    }
    return raw, err
}

// GetObserverList returns the list of observer nodes' ID according to the groupID
func (gc *Client) GetObserverList(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getObserverList", groupID)
    if err != nil {
        return nil, err
    }
    return raw, err
}

// TODO
// process the json result of GetConsensusStatus
// type consensusStatus struct {

// }

// type commonStatus struct {
//     AccountType uint           
//     AllowFutureBlocks bool    
//     CfgErr bool
//     connectedNodes uint
//     consensusedBlockNumber uint
//     currentView uint
//     groupId uint
//     highestblockHash string
//     highestblockNumber uint
//     leaderFailed bool
//     max_faulty_leader uint
// }

// GetConsensusStatus returns the status information about the consensus algorithm on a specific groupID
func (gc *Client) GetConsensusStatus(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getConsensusStatus", groupID)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetSyncStatus returns the synchronization status of the group
func (gc *Client) GetSyncStatus(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getSyncStatus", groupID)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetPeers returns the information of the connected peers
func (gc *Client) GetPeers(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getPeers", groupID)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetGroupPeers returns the nodes and the overser nodes list on a specific group
func (gc *Client) GetGroupPeers(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getGroupPeers", groupID)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetNodeIDList returns the ID information of the connected peers and itself
func (gc *Client) GetNodeIDList(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getNodeIDList", groupID)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// getGroupList returns the groupID that the node belongs to 
func (gc *Client) GetGroupList(ctx context.Context) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getGroupList")
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetBlockByHash returns the block information according to the given block hash
func (gc *Client) GetBlockByHash(ctx context.Context, groupID *big.Int, bhash string, includetx bool) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getBlockByHash", groupID, bhash, includetx)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetBlockByNumber returns the block information according to the given block number(hex format)
func (gc *Client) GetBlockByNumber(ctx context.Context, groupID *big.Int, bnum string, includetx bool) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getBlockByNumber", groupID, bnum, includetx)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetBlockHashByNumber returns the block hash according to the given block number
func (gc *Client) GetBlockHashByNumber(ctx context.Context, groupID *big.Int, bnum string) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getBlockHashByNumber", groupID, bnum)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetTransactionByHash returns the transaction information according to the given transaction hash
func (gc *Client) GetTransactionByHash(ctx context.Context, groupID *big.Int, txhash string) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getTransactionByHash", groupID, txhash)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// getTransactionByBlockHashAndIndex returns the transaction information according to 
// the given block hash and transaction index
func (gc *Client) GetTransactionByBlockHashAndIndex(ctx context.Context, groupID *big.Int, bhash string, txindex string) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getTransactionByBlockHashAndIndex", groupID, bhash, txindex)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetTransactionByBlockNumberAndIndex returns the transaction information according to 
// the given block number and transaction index
func (gc *Client) GetTransactionByBlockNumberAndIndex(ctx context.Context, groupID *big.Int, bnum string, txindex string) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getTransactionByBlockNumberAndIndex", groupID, bnum, txindex)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetTransactionReceipt returns the transaction receipt according to the given transaction hash
func (gc *Client) GetTransactionReceipt(ctx context.Context, groupID *big.Int, txhash string) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getTransactionReceipt", groupID, txhash)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetPendingTransactions returns information of the pending transactions
func (gc *Client) GetPendingTransactions(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getPendingTransactions", groupID)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetPendingTxSize returns amount of the pending transactions
func (gc *Client) GetPendingTxSize(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getPendingTxSize", groupID)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetCode returns the contract code according to the contract address
func (gc *Client) GetCode(ctx context.Context, groupID *big.Int, addr string) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getCode", groupID, addr)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetTotalTransactionCount returns the totoal amount of transactions and the block height at present
func (gc *Client) GetTotalTransactionCount(ctx context.Context, groupID *big.Int) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getTotalTransactionCount", groupID)
    if err != nil {
        return nil, err
    }
    return raw,err
}

// GetSystemConfigByKey returns value according to the key(only tx_count_limit, tx_gas_limit could work)
func (gc *Client) GetSystemConfigByKey(ctx context.Context, groupID *big.Int, findkey string) ([]byte, error) {
    var raw json.RawMessage
    err := gc.c.CallContext(ctx, &raw, "getSystemConfigByKey", groupID, findkey)
    if err != nil {
        return nil, err
    }
    return raw,err
}