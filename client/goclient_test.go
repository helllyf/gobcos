package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/KasperLiu/gobcos/accounts/abi/bind"
	"github.com/KasperLiu/gobcos/common"
	"github.com/KasperLiu/gobcos/common/hexutil"
	"github.com/KasperLiu/gobcos/contract/Store"
	"github.com/KasperLiu/gobcos/contract/router"
	"github.com/KasperLiu/gobcos/crypto"
	"log"
	"testing"
)

func GetClient(t *testing.T) (*Client) {
	// RPC API
	groupID := uint(1)
	c, err := Dial("http://192.168.1.131:8545", groupID)  // change to your RPC and groupID
	if err != nil {
		t.Fatalf("can not dial to the RPC API: %v", err)
	}
	return c
}

func TestClientVersion(t *testing.T) {
	c := GetClient(t)

	cv, err := c.GetClientVersion(context.Background())
	if err != nil {
		t.Fatalf("client version not found: %v", err)
	}

	t.Logf("client version:\n%s", cv)
}

func TestBlockNumber(t *testing.T) {
    c := GetClient(t)
	// cannot use big.NewInt to construct json request
	// TODO: analysis the ethereum's big.NewInt
	bn, err := c.GetBlockNumber(context.Background())
	if err != nil {
		t.Fatalf("block number not found: %v", err)
	}

	t.Logf("latest block number: \n%s", bn)
}

func TestPBFTView(t *testing.T) {
	c := GetClient(t)
	pv, err := c.GetPBFTView(context.Background())
	if err != nil {
		t.Fatalf("PBFT view not found: %v", err)
	}

	t.Logf("PBFT view: \n%s", pv)
}

func TestCreateAccount(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
}

func TestRouter(t *testing.T) {
	client := GetClient(t)
	privateKey, _ := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
	_address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(_address)
	auth := bind.NewKeyedTransactor(privateKey)
	address := common.HexToAddress("0xC1Be3A354D20696B5c91626322835C81F2000c82")
	instance, _ := router.NewRouter(address, client)
	//auth.From = crypto.PubkeyToAddress(*publicKeyECDSA)
	tx,err := instance.InsertAuto(auth)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("transaction hash: ", tx.Hash().Hex())
	sender := common.HexToAddress("0xE280029a7867BA5C9154434886c241775ea87e53")
	opts := &bind.CallOpts{From: sender}
	res,err := instance.Show(opts,sender)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(res.Hex())
}

func TestDeploy(t *testing.T) {
	client := GetClient(t)
	//privateKey, err := crypto.HexToECDSA("input your privateKey in hex without \"0x\"") // 145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58

	privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
	auth := bind.NewKeyedTransactor(privateKey) // input your privateKey
	input := "Store deployment 1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	receipt, err := bind.WaitMined(context.Background(),client,tx)
	if err != nil {
		log.Fatalf("tx mining error:%v\n", err)
	}
	fmt.Println(receipt.ContractAddress.Hex())
	address1,tx1,instance1,err1 := router.DeployRouter(auth,client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex())  // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())
	_ = instance
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("contract address: ", address1.Hex())  // the address should be saved
	fmt.Println("transaction hash: ", tx1.Hash().Hex())
	_ = instance1
}

func TestLoad(t *testing.T) {
	client := GetClient(t)
	address := common.HexToAddress("0xEC129c56486AD3504A8922F8dc45A07941f13e05")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")
	_ = instance
}

func TestRead(t *testing.T) {
	client := GetClient(t)
	address := common.HexToAddress("0xEC129c56486AD3504A8922F8dc45A07941f13e05")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	opts := &bind.CallOpts{From: common.HexToAddress("0xe280029a7867ba5c9154434886c241775ea87e53")} //0xFbb18d54e9Ee57529cda8c7c52242EFE879f064F
	version, err := instance.Version(opts)
	key := [32]byte{}
	copy(key[:], []byte("foo"))

	value,err :=  instance.Items(opts,key)
	if err != nil {
		log.Fatal(err)
	}
	resvalue := value[:]
	fmt.Println(common.Bytes2Hex(resvalue))

	fmt.Println("version :", version) // "Store deployment 1.0"
}

func TestDemo(t *testing.T)  {
	// load the contract
	client := GetClient(t)
	address := common.HexToAddress("0xEC129c56486AD3504A8922F8dc45A07941f13e05")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5") // 145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func TestWrite(t *testing.T) {
	client := GetClient(t)
	address := common.HexToAddress("0xEC129c56486AD3504A8922F8dc45A07941f13e05")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar111"))

	privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5") // 145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
// func TestBlockLimit(t *testing.T) {
//     c := GetClient(t)
// 	// cannot use big.NewInt to construct json request
// 	// TODO: analysis the ethereum's big.NewInt
// 	bl, err := c.GetBlockLimit(context.Background())
// 	if err != nil {
// 		t.Fatalf("blockLimit not found: %v", err)
// 	}

// 	t.Logf("latest blockLimit: \n%s", bl)
// }

// func TestGroupID(t *testing.T) {
//     c := GetClient(t)
// 	// cannot use big.NewInt to construct json request
// 	// TODO: analysis the ethereum's big.NewInt
// 	groupid := c.GetGroupID()
// 	t.Logf("current groupID: \n%s", groupid)
// }

// func TestChainID(t *testing.T) {
//     c := GetClient(t)
// 	// cannot use big.NewInt to construct json request
// 	// TODO: analysis the ethereum's big.NewInt
// 	chainid, err := c.GetChainID(context.Background())
// 	if err != nil {
// 		t.Fatalf("Chain ID not found: %v", err)
// 	}
// 	t.Logf("Chain ID: \n%s", chainid)
// }

// func TestSealerList(t *testing.T) {
// 	c := GetClient(t)
// 	sl, err := c.GetSealerList(context.Background())
// 	if err != nil {
// 		t.Fatalf("sealer list not found: %v", err)
// 	}

// 	t.Logf("sealer list:\n%s", sl)
// }

// func TestObserverList(t *testing.T) {
// 	c := GetClient(t)
// 	ol, err := c.GetObserverList(context.Background())
// 	if err != nil {
// 		t.Fatalf("observer list not found: %v", err)
// 	}

// 	t.Logf("observer list:\n%s", ol)
// }

// func TestConsensusStatus(t *testing.T) {
// 	c := GetClient(t)
// 	status, err := c.GetConsensusStatus(context.Background())
// 	if err != nil {
// 		t.Fatalf("consensus status not found: %v", err)
// 	}

// 	t.Logf("consensus status:\n%s", status)
// }

// func TestSyncStatus(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetSyncStatus(context.Background())
// 	if err != nil {
// 		t.Fatalf("synchronization status not found: %v", err)
// 	}

// 	t.Logf("synchronization Status:\n%s", raw)
// }

// func TestPeers(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetPeers(context.Background())
// 	if err != nil {
// 		t.Fatalf("peers not found: %v", err)
// 	}

// 	t.Logf("peers:\n%s", raw)
// }

// func TestGroupPeers(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetGroupPeers(context.Background())
// 	if err != nil {
// 		t.Fatalf("group peers not found: %v", err)
// 	}

// 	t.Logf("group peers:\n%s", raw)
// }

// func TestNodeIDList(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetNodeIDList(context.Background())
// 	if err != nil {
// 		t.Fatalf("nodeID list not found: %v", err)
// 	}

// 	t.Logf("nodeID list:\n %s", raw)
// }

// func TestGroupList(t *testing.T) {
// 	c := GetClient(t)
// 	raw, err := c.GetGroupList(context.Background())
// 	if err != nil {
// 		t.Fatalf("group list not found: %v", err)
// 	}

// 	t.Logf("group list:\n%s", raw)
// }

// func TestBlockByHash(t *testing.T) {
// 	c := GetClient(t)
	
// 	bhash := "0xc0b21d064b97bafda716e07785fe8bb20cc23506bb980f12c7f7a4f4ef50ce30"
// 	includeTx := false
// 	raw, err := c.GetBlockByHash(context.Background(), bhash, includeTx)
// 	if err != nil {
// 		t.Fatalf("block not found: %v", err)
// 	}

// 	t.Logf("block by hash:\n%s", raw)
// }

// func TestBlockByNumber(t *testing.T) {
// 	c := GetClient(t)
	
// 	bnum := "0x1"
// 	includeTx := true
// 	raw, err := c.GetBlockByNumber(context.Background(), bnum, includeTx)
// 	if err != nil {
// 		t.Fatalf("block not found: %v", err)
// 	}

// 	t.Logf("block by number:\n%s", raw)
// }

// func TestBlockHashByNumber(t *testing.T) {
// 	c := GetClient(t)
	
// 	bnum := "0x1"
// 	raw, err := c.GetBlockHashByNumber(context.Background(), bnum)
// 	if err != nil {
// 		t.Fatalf("block hash not found: %v", err)
// 	}

// 	t.Logf("block hash by number:\n%s", raw)
// }

// func TestTransactionByHash(t *testing.T) {
// 	c := GetClient(t)
	
// 	txhash := "0xed51827558939e8d103cbf8f6ff37f8a99582f09afa29e5636d0e54a073d0893"
// 	raw, err := c.GetTransactionByHash(context.Background(), txhash)
// 	if err != nil {
// 		t.Fatalf("transaction not found: %v", err)
// 	}

// 	t.Logf("transaction by hash:\n%s", raw)
// }

// func TestTransactionByBlockHashAndIndex(t *testing.T) {
// 	c := GetClient(t)
	
// 	bhash := "0xc0b21d064b97bafda716e07785fe8bb20cc23506bb980f12c7f7a4f4ef50ce30"
// 	txindex := "0x0"
// 	raw, err := c.GetTransactionByBlockHashAndIndex(context.Background(), bhash, txindex)
// 	if err != nil {
// 		t.Fatalf("transaction not found: %v", err)
// 	}

// 	t.Logf("transaction by block hash and transaction index:\n%s", raw)
// }

// func TestTransactionByBlockNumberAndIndex(t *testing.T) {
// 	c := GetClient(t)
	
// 	bnum := "0x1"
// 	txindex := "0x0"
// 	raw, err := c.GetTransactionByBlockNumberAndIndex(context.Background(), bnum, txindex)
// 	if err != nil {
// 		t.Fatalf("transaction not found: %v", err)
// 	}

// 	t.Logf("transaction by block number and transaction index:\n%s", raw)
// }

// func TestTransactionReceipt(t *testing.T) {
// 	c := GetClient(t)
	
// 	txhash := "0xed51827558939e8d103cbf8f6ff37f8a99582f09afa29e5636d0e54a073d0893"
// 	raw, err := c.GetTransactionReceipt(context.Background(), txhash)
// 	if err != nil {
// 		t.Fatalf("transaction receipt not found: %v", err)
// 	}

// 	t.Logf("transaction receipt by transaction hash:\n%s", raw)
// }

// func TestContractAddress(t *testing.T) {
// 	c := GetClient(t)
// 	txhash := "0x4a2a4d878318a83491383d29d6550c088bdcf692e3055b060342dcd85177c621"
// 	ca, err := c.GetContractAddress(context.Background(), txhash)
// 	if err != nil {
// 		t.Fatalf("ContractAddress not found: %v", err)
// 	}

// 	t.Logf("ContractAddress: \n%s", ca.String())
// }

// func TestPendingTransactions(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetPendingTransactions(context.Background())
// 	if err != nil {
// 		t.Fatalf("pending transactions not found: %v", err)
// 	}

// 	t.Logf("pending transactions:\n%s", raw)
// }

// func TestPendingTxSize(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetPendingTxSize(context.Background())
// 	if err != nil {
// 		t.Fatalf("pending transactions not found: %v", err)
// 	}

// 	t.Logf("the amount of the pending transactions:\n%s", raw)
// }

// func TestGetCode(t *testing.T) {
// 	c := GetClient(t)
	
// 	addr := "0x27c1b5d9fe3ab035c2e9db7199d4beb139e12292"
// 	raw, err := c.GetCode(context.Background(), addr)
// 	if err != nil {
// 		t.Fatalf("contract not found: %v", err)
// 	}

// 	t.Logf("the contract code:\n%s", raw)
// }

// func TestTotalTransactionCount(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetTotalTransactionCount(context.Background())
// 	if err != nil {
// 		t.Fatalf("transactions not found: %v", err)
// 	}

// 	t.Logf("the totoal transactions and present block height:\n%s", raw)
// }

// func TestSystemConfigByKey(t *testing.T) {
// 	c := GetClient(t)
	
// 	findkey := "tx_count_limit"
// 	raw, err := c.GetSystemConfigByKey(context.Background(), findkey)
// 	if err != nil {
// 		t.Fatalf("the value not found: %v", err)
// 	}

// 	t.Logf("the value got by the key:\n%s", raw)
// }
