package netHelper

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"sync"
)

/**
* @Author: liujunren
* @Date: 2022/3/8 9:21
 */
var (
	file_proto_msgTypes    = make([]protoimpl.MessageInfo, 6)
	file_proto_rawDescOnce sync.Once
)

func NewMessage(src interface{}) proto.Message {
	marshal, _ := json.Marshal(src)
	s := string(marshal)
	return &Message{Data:&s }
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Data          *string `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

// Deprecated: Use CheckPwdRes.ProtoReflect.Descriptor instead.
//func (*Message) Descriptor() ([]byte, []int) {
//	return file_proto_rawDescGZIP(), []int{5}
//}

//func file_proto_rawDescGZIP() []byte {
//	file_proto_rawDescOnce.Do(func() {
//		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
//	})
//	return file_admin_proto_rawDescData
//}
//
//func (x *Message) GetCode() int32 {
//	if x != nil {
//		return x.Code
//	}
//	return 0
//}
//
//func (x *Message) GetMsg() string {
//	if x != nil {
//		return x.Msg
//	}
//	return ""
//}
//
//func (x *Message) GetData() *anypb.Any {
//	if x != nil {
//		return x.Data
//	}
//	return nil
//}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
