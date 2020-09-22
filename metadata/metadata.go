package metadata

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/metadata"
)

type UserAgent struct {
	AccountID string `json:"account_id"`
	UID       string `json:"uid"`
	UserType  int    `json:"user_type"` // 1:平台管理员，2：机构内管理员，2：普通用户

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
