// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package example

import (
	"context"
	"encoding/json"
)

type ClientExampleRPC struct {
	*ClientJsonRPC
}

type retExampleRPCTest func(ret1 int, ret2 string, err error)

func (cli *ClientExampleRPC) ReqTest(ret retExampleRPCTest, arg0 int, arg1 string, opts ...interface{}) (request baseJsonRPC) {

	request = baseJsonRPC{
		Method: "examplerpc.test",
		Params: requestExampleRPCTest{
			Arg0: arg0,
			Arg1: arg1,
			Opts: opts,
		},
		Version: Version,
	}
	var err error
	var response responseExampleRPCTest

	if ret != nil {
		request.retHandler = func(jsonrpcResponse baseJsonRPC) {
			if jsonrpcResponse.Error != nil {
				err = cli.errorDecoder(jsonrpcResponse.Error)
				ret(response.Ret1, response.Ret2, err)
				return
			}
			err = json.Unmarshal(jsonrpcResponse.Result, &response)
			ret(response.Ret1, response.Ret2, err)
		}
		request.ID = []byte("\"" + uuid.NewV4().String() + "\"")
	}
	return
}

func (cli *ClientExampleRPC) Test(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (ret1 int, ret2 string, err error) {

	retHandler := func(_ret1 int, _ret2 string, _err error) {
		ret1 = _ret1
		ret2 = _ret2
		err = _err
	}
	if blockErr := cli.Batch(ctx, cli.ReqTest(retHandler, arg0, arg1, opts...)); blockErr != nil {
		err = blockErr
		return
	}
	return
}
