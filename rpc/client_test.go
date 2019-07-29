package rpc

import (
	"testing"
)

func GetClient(t *testing.T) (*rpcClient) {
	client, err := Dial("http://localhost:8545")
	if err != nil{
		t.Fatalf("can't dial to the json-rpc API: %v", err)
	}
	return &client
}

func TestCall(t *testing.T) {
	client := GetClient(t)
	var cVersion interface{}
	err := client.Call(&cVersion, "getClientVersion")
	if err != nil {
		t.Fatalf("invoke json-rpc API error occur: %v", err)
	}
	t.Logf("client version: %s", cVersion)
}

func TestCallJsonResponse(t *testing.T) {
	client := GetClient(t)
	resp, err := client.CallJsonResponse("getClientVersion")
    if err != nil {
		t.Fatalf("invoke json-rpc API error occur: %v", err)
	}
    t.Logf("client version: \n %s", resp)
}