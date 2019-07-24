package client

import (
	"context"
	"testing"
	"math/big"
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
	defer c.Close()
	
	cv, err := c.GetClientVersion(context.Background())
	if err != nil {
		t.Fatalf("client version not found: %v", err)
	}

	t.Log(cv)
}

func TestBlockNumber(t *testing.T) {
    c := GetClient(t)
	defer c.Close()
	
	groupID := big.NewInt(1)
	bn, err := c.GetBlockNumber(context.Background(), groupID)
	if err != nil {
		t.Fatalf("block number not found: %v", err)
	}

	t.Logf("latest block number: %s", bn)
}