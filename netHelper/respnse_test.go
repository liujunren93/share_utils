package netHelper

import (
	"github.com/shareChina/utils/status"
	"google.golang.org/protobuf/runtime/protoimpl"
	"testing"
)

type DefaultRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *DefaultRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DefaultRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DefaultRes) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func TestRpcResponse(t *testing.T) {
	var aa *DefaultRes
	RpcResponse(&(*aa), status.StatusOK, "adsa", "as")
}
