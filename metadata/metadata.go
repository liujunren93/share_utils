package metadata

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/metadata"
)

type UserAgent struct {
	CompanyID  uint `json:"company_id"`  // 机构id
	AccountID  uint `json:"account_id"`  // 用户账户id
	UID        uint `json:"uid"`         // 当前用户ID 服务用户id
	UserType   int8 `json:"user_type"`   // 1:平台管理员，2：机构内管理员，2：普通用户
	IsRoot     bool `json:"is_root"`     // 是否是超管
	ClientType int8 `json:"client_type"` // 1:平台，2：机构后台，2：客户端
	ClientID   int8 `json:"client_id"` // 客户端id
}

//获取
func GetValue(ctx context.Context, key string) ([]string, bool) {
	incomingContext, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ok
	}
	strings, ok := incomingContext[key]
	return strings, ok
}

//获取
func GetAll(ctx context.Context) (metadata.MD, bool) {
	return metadata.FromIncomingContext(ctx)

}

func GetUA(ctx context.Context) (*UserAgent, bool) {
	value, ok := GetValue(ctx, "ua")
	if !ok {
		return nil, ok
	}
	s := value[0]
	var data UserAgent
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		return nil, false
	}
	return &data, true

}

func SetUA(ctx context.Context, agent *UserAgent) context.Context {
	marshal, _ := json.Marshal(agent)
	pairs := metadata.Pairs("ua", string(marshal))
	return metadata.NewOutgoingContext(ctx, pairs)
}
