package httpHandler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/shareChina/utils/httpHelper"
	"time"
)

type rpcRequest struct {
	service  string
	endpoint string
	method   string
	address  string
	timeout  int
	request  interface{}
}

type RpcResult struct {
	Code int32             `json:"code"`
	Msg  httpHelper.Option `json:"msg"`
	Data interface{}       `json:"data"`
}

func NewRPCRequest(service, endpoint, method, address string, request interface{}) (*rpcRequest, error) {
	if service == "" {
		return nil, errors.New("service cannot be empty")
	}
	if endpoint == "" {
		return nil, errors.New("endpoint cannot be empty")
	}
	if method == "" {
		return nil, errors.New("method cannot be empty")
	}
	if request == nil {
		return nil, errors.New("request cannot be nil")
	}
	return &rpcRequest{
		service:  service,
		endpoint: endpoint,
		method:   method,
		address:  address,
		request:  request,
	}, nil
}

func (r *rpcRequest) RPC(ctx context.Context) (res *RpcResult, err error) {

	request := (*cmd.DefaultOptions().Client).NewRequest(r.service, r.endpoint, r.request, client.WithContentType("application/json"))

	var opts []client.CallOption
	if len(r.address) > 0 {
		opts = append(opts, client.WithAddress(r.address))
	}
	if r.timeout > 0 {
		opts = append(opts, client.WithRequestTimeout(time.Duration(r.timeout)*time.Second))
	}
	var response json.RawMessage
	err = (*cmd.DefaultOptions().Client).Call(ctx, request, &response, opts...)
	if err != nil {
		return
	}
	marshalJSON, err := response.MarshalJSON()
	if err != nil {
		return
	}
	err = json.Unmarshal(marshalJSON, res)

	return
}
