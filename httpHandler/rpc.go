package httpHandler

import (
	"errors"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config/cmd"
	"net/http"
	"strconv"
)

type rpcRequest struct {
	service  string
	endpoint string
	method   string
	address  string
	request  interface{}
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

func (r *rpcRequest) RPC(req *http.Request) {
	request := (*cmd.DefaultOptions().Client).NewRequest(r.service, r.endpoint, r.request, client.WithContentType("application/json"))
	timeout, _ := strconv.Atoi(req.Header.Get("Timeout"))
	var opts []client.CallOption
	if len(r.address) > 0 {
		opts = append(opts, client.WithAddress(r.address))
	}
	err = (*cmd.DefaultOptions().Client).Call(ctx, req, &response, opts...)

}
