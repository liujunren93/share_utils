package sts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

type STS struct {
	endpoint string
	client   *sts.Client
}

func NewSTS() (*STS, error) {
	client, err := sts.NewClientWithAccessKey("cn-shanghai", "LTAI5tDPQZYtaCVJ7XtK8noS", "G5LuLDDHqykuQnOmfPTXDabgWcMzRm")
	if err != nil {
		return nil, err
	}
	return &STS{
		endpoint: "",
		client:   client,
	}, nil
}
func (s *STS)Token() (*sts.AssumeRoleResponse,error) {
	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见API参考。
	request.RoleArn ="acs:ram::1855340179767769:role/osssts"
	request.RoleSessionName ="safdsafdsafd"

	//发起请求，并得到响应。
	return s.client.AssumeRole(request)

}
