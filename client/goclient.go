// client provides the RPC methods for FISCO BCOS
package client

import (
    "errors"
    
    "gobcos/rpc"
)

// client defines typed wrapper for the FISCO BCOS RPC API
// refer the Ethereum
type Client struct {
    c rpc.RPCClient
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
    client, err := rpc.Dial("http://localhost:8545")
	if err != nil{
        errs := errors.New("can't dial to the json-rpc API: " + err.Error())
        return nil, errs
	}
	return &Client{c:client}, err
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
func (gc *Client) GetClientVersion() ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getClientVersion")
    if err != nil {
		return nil, err
	}
    return resp,err
}

// GetClientVersion returns the data struct wih the version of FISCO BCOS running on the nodes.
// func (gc *Client) GetClientVersionDS() (*clientVersion, error) {
//     var cv *clientVersion
//     resp, err := gc.c.Call(cv, "getClientVersion")
//     if err != nil {
// 		return nil, err
// 	}
//     return resp,err
// }

// GetBlockNumber returns the latest block height(hex format) on a given groupID.
func (gc *Client) GetBlockNumber(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getBlockNumber", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetPBFTView returns the latest PBFT view(hex format) of the specific group and it will returns a wrong sentence
// if the consensus algorithm is not the PBFT.
func (gc *Client) GetPBFTView(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getPbftView", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err

    // TODO
    // Raft consensus
}

// GetSealerList returns the list of consensus nodes' ID according to the groupID
func (gc *Client) GetSealerList(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getSealerList", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetObserverList returns the list of observer nodes' ID according to the groupID
func (gc *Client) GetObserverList(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getObserverList", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
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
func (gc *Client) GetConsensusStatus(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getConsensusStatus", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetSyncStatus returns the synchronization status of the group
func (gc *Client) GetSyncStatus(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getSyncStatus", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetPeers returns the information of the connected peers
func (gc *Client) GetPeers(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getPeers", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetGroupPeers returns the nodes and the overser nodes list on a specific group
func (gc *Client) GetGroupPeers(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getGroupPeers", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetNodeIDList returns the ID information of the connected peers and itself
func (gc *Client) GetNodeIDList(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getNodeIDList", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// getGroupList returns the groupID that the node belongs to 
func (gc *Client) GetGroupList() ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getGroupList")
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetBlockByHash returns the block information according to the given block hash
func (gc *Client) GetBlockByHash(groupID uint, bhash string, includetx bool) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getBlockByHash", groupID, bhash, includetx)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetBlockByNumber returns the block information according to the given block number(hex format)
func (gc *Client) GetBlockByNumber(groupID uint, bnum string, includetx bool) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getBlockByNumber", groupID, bnum, includetx)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetBlockHashByNumber returns the block hash according to the given block number
func (gc *Client) GetBlockHashByNumber(groupID uint, bnum string) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getBlockHashByNumber", groupID, bnum)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetTransactionByHash returns the transaction information according to the given transaction hash
func (gc *Client) GetTransactionByHash(groupID uint, txhash string) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getTransactionByHash", groupID, txhash)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// getTransactionByBlockHashAndIndex returns the transaction information according to 
// the given block hash and transaction index
func (gc *Client) GetTransactionByBlockHashAndIndex(groupID uint, bhash string, txindex string) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getTransactionByBlockHashAndIndex", groupID, bhash, txindex)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetTransactionByBlockNumberAndIndex returns the transaction information according to 
// the given block number and transaction index
func (gc *Client) GetTransactionByBlockNumberAndIndex(groupID uint, bnum string, txindex string) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getTransactionByBlockNumberAndIndex", groupID, bnum, txindex)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetTransactionReceipt returns the transaction receipt according to the given transaction hash
func (gc *Client) GetTransactionReceipt(groupID uint, txhash string) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getTransactionReceipt", groupID, txhash)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetPendingTransactions returns information of the pending transactions
func (gc *Client) GetPendingTransactions(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getPendingTransactions", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetPendingTxSize returns amount of the pending transactions
func (gc *Client) GetPendingTxSize(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getPendingTxSize", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetCode returns the contract code according to the contract address
func (gc *Client) GetCode(groupID uint, addr string) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getCode", groupID, addr)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetTotalTransactionCount returns the totoal amount of transactions and the block height at present
func (gc *Client) GetTotalTransactionCount(groupID uint) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getTotalTransactionCount", groupID)
    if err != nil {
        return nil, err
    }
    return resp, err
}

// GetSystemConfigByKey returns value according to the key(only tx_count_limit, tx_gas_limit could work)
func (gc *Client) GetSystemConfigByKey(groupID uint, findkey string) ([]byte, error) {
    resp, err := gc.c.CallJsonResponse("getSystemConfigByKey", groupID, findkey)
    if err != nil {
        return nil, err
    }
    return resp, err
}