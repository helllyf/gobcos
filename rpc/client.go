package rpc

import (
	"encoding/json"

	"github.com/ybbus/jsonrpc"
)

const (
	contentType = "application/json"
)

type RPCClient struct {
	rpc jsonrpc.RPCClient
}

// Dial to the RPC server to get the information
func Dial(endpoint string) (RPCClient, error){
	r := jsonrpc.NewClientWithOpts(endpoint, &jsonrpc.RPCClientOpts{
		CustomHeaders: map[string]string{
			"Content-Type": contentType,
			"Accept"      : contentType,
		},
	})

	// test the request
	_, err := r.Call("getClientVersion")

	return RPCClient{rpc:r}, err
}

// Call the RPC method and return the Json result according to the param
func (r *RPCClient) Call(result interface{}, method string, args ...interface{}) error {
	switch response, err := r.rpc.Call(method, args...); {
	case err != nil:
		// error handling goes here e.g. network / http error
		return err
	case response.Error != nil:
		// rpc error handling goes here
		// check response.Error.Code, response.Error.Message and optional response.Error.Data
		return response.Error
	default:
		// some error on json umarshal level or json result field was null
		return response.GetObject(&result)
	}
}

// Call the RPC method and return the Json result directly
func (r *RPCClient) CallJsonResponse(method string, args ...interface{}) ([]byte ,error) {
	switch response, err := r.rpc.Call(method, args...); {
	case err != nil:
		// error handling goes here e.g. network / http error
		return nil, err
	case response.Error != nil:
		// rpc error handling goes here
		// check response.Error.Code, response.Error.Message and optional response.Error.Data
		return nil, response.Error
	default:
		// some error on json umarshal level or json result field was null
		js, err := json.MarshalIndent(response.Result, "", "\t")
		if err != nil {
			return nil, err
		}
		return js, err
	}
}