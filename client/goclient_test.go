package client

import (
	"testing"
)

func GetClient(t *testing.T) (*Client) {
	// RPC API
	c, err := Dial("http://localhost:8545")
	if err != nil {
		t.Fatalf("can not dial to the RPC API: %v", err)
	}
	return c
}

func TestClientVersion(t *testing.T) {
	c := GetClient(t)

	cv, err := c.GetClientVersion()
	if err != nil {
		t.Fatalf("client version not found: %v", err)
	}

	t.Logf("client version:\n%s", cv)
}

func TestBlockNumber(t *testing.T) {
    c := GetClient(t)
	// cannot use big.NewInt to construct json request
	// TODO: analysis the ethereum's big.NewInt
	var groupID uint
	groupID = 1
	bn, err := c.GetBlockNumber(groupID)
	if err != nil {
		t.Fatalf("block number not found: %v", err)
	}

	t.Logf("latest block number: \n%s", bn)
}

// func TestPBFTView(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	pv, err := c.GetPBFTView(groupID)
// 	if err != nil {
// 		t.Fatalf("PBFT view not found: %v", err)
// 	}

// 	t.Logf("PBFT view: \n%s", pv)
// }

// func TestSealerList(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	sl, err := c.GetSealerList(groupID)
// 	if err != nil {
// 		t.Fatalf("sealer list not found: %v", err)
// 	}

// 	t.Logf("sealer list:\n%s", sl)
// }

// func TestObserverList(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	ol, err := c.GetObserverList(groupID)
// 	if err != nil {
// 		t.Fatalf("observer list not found: %v", err)
// 	}

// 	t.Logf("observer list:\n%s", ol)
// }

// func TestConsensusStatus(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	status, err := c.GetConsensusStatus(groupID)
// 	if err != nil {
// 		t.Fatalf("consensus status not found: %v", err)
// 	}

// 	t.Logf("consensus status:\n%s", status)
// }

// func TestSyncStatus(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	raw, err := c.GetSyncStatus(groupID)
// 	if err != nil {
// 		t.Fatalf("synchronization status not found: %v", err)
// 	}

// 	t.Logf("synchronization Status:\n%s", raw)
// }

// func TestPeers(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	raw, err := c.GetPeers(groupID)
// 	if err != nil {
// 		t.Fatalf("peers not found: %v", err)
// 	}

// 	t.Logf("peers:\n%s", raw)
// }

// func TestGroupPeers(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	raw, err := c.GetGroupPeers(groupID)
// 	if err != nil {
// 		t.Fatalf("group peers not found: %v", err)
// 	}

// 	t.Logf("group peers:\n%s", raw)
// }

// func TestNodeIDList(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	raw, err := c.GetNodeIDList(groupID)
// 	if err != nil {
// 		t.Fatalf("nodeID list not found: %v", err)
// 	}

// 	t.Logf("nodeID list:\n %s", raw)
// }

// func TestGroupList(t *testing.T) {
// 	c := GetClient(t)
// 	raw, err := c.GetGroupList()
// 	if err != nil {
// 		t.Fatalf("group list not found: %v", err)
// 	}

// 	t.Logf("group list:\n%s", raw)
// }

// func TestBlockByHash(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	bhash := "0xc0b21d064b97bafda716e07785fe8bb20cc23506bb980f12c7f7a4f4ef50ce30"
// 	includeTx := false
// 	raw, err := c.GetBlockByHash(groupID, bhash, includeTx)
// 	if err != nil {
// 		t.Fatalf("block not found: %v", err)
// 	}

// 	t.Logf("block by hash:\n%s", raw)
// }

// func TestBlockByNumber(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	bnum := "0x1"
// 	includeTx := true
// 	raw, err := c.GetBlockByNumber(groupID, bnum, includeTx)
// 	if err != nil {
// 		t.Fatalf("block not found: %v", err)
// 	}

// 	t.Logf("block by number:\n%s", raw)
// }

// func TestBlockHashByNumber(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	bnum := "0x1"
// 	raw, err := c.GetBlockHashByNumber(groupID, bnum)
// 	if err != nil {
// 		t.Fatalf("block hash not found: %v", err)
// 	}

// 	t.Logf("block hash by number:\n%s", raw)
// }

// func TestTransactionByHash(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	txhash := "0xed51827558939e8d103cbf8f6ff37f8a99582f09afa29e5636d0e54a073d0893"
// 	raw, err := c.GetTransactionByHash(groupID, txhash)
// 	if err != nil {
// 		t.Fatalf("transaction not found: %v", err)
// 	}

// 	t.Logf("transaction by hash:\n%s", raw)
// }

// func TestTransactionByBlockHashAndIndex(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	bhash := "0xc0b21d064b97bafda716e07785fe8bb20cc23506bb980f12c7f7a4f4ef50ce30"
// 	txindex := "0x0"
// 	raw, err := c.GetTransactionByBlockHashAndIndex(groupID, bhash, txindex)
// 	if err != nil {
// 		t.Fatalf("transaction not found: %v", err)
// 	}

// 	t.Logf("transaction by block hash and transaction index:\n%s", raw)
// }

// func TestTransactionByBlockNumberAndIndex(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	bnum := "0x1"
// 	txindex := "0x0"
// 	raw, err := c.GetTransactionByBlockNumberAndIndex(groupID, bnum, txindex)
// 	if err != nil {
// 		t.Fatalf("transaction not found: %v", err)
// 	}

// 	t.Logf("transaction by block number and transaction index:\n%s", raw)
// }

// func TestTransactionReceipt(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	txhash := "0xed51827558939e8d103cbf8f6ff37f8a99582f09afa29e5636d0e54a073d0893"
// 	raw, err := c.GetTransactionReceipt(groupID, txhash)
// 	if err != nil {
// 		t.Fatalf("transaction receipt not found: %v", err)
// 	}

// 	t.Logf("transaction receipt by transaction hash:\n%s", raw)
// }

// func TestPendingTransactions(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	raw, err := c.GetPendingTransactions(groupID)
// 	if err != nil {
// 		t.Fatalf("pending transactions not found: %v", err)
// 	}

// 	t.Logf("pending transactions:\n%s", raw)
// }

// func TestPendingTxSize(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	raw, err := c.GetPendingTxSize(groupID)
// 	if err != nil {
// 		t.Fatalf("pending transactions not found: %v", err)
// 	}

// 	t.Logf("the amount of the pending transactions:\n%s", raw)
// }

// func TestGetCode(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	addr := "0x27c1b5d9fe3ab035c2e9db7199d4beb139e12292"
// 	raw, err := c.GetCode(groupID, addr)
// 	if err != nil {
// 		t.Fatalf("contract not found: %v", err)
// 	}

// 	t.Logf("the contract code:\n%s", raw)
// }

// func TestTotalTransactionCount(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	raw, err := c.GetTotalTransactionCount(groupID)
// 	if err != nil {
// 		t.Fatalf("transactions not found: %v", err)
// 	}

// 	t.Logf("the totoal transactions and present block height:\n%s", raw)
// }

// func TestSystemConfigByKey(t *testing.T) {
// 	c := GetClient(t)
// 	var groupID uint
// 	groupID = 1
// 	findkey := "tx_count_limit"
// 	raw, err := c.GetSystemConfigByKey(groupID, findkey)
// 	if err != nil {
// 		t.Fatalf("the value not found: %v", err)
// 	}

// 	t.Logf("the value got by the key:\n%s", raw)
// }