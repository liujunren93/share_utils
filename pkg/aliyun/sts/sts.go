package sts

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/liujunren93/share_utils/helper"
)

type STS struct {
	endpoint string
	client   *sts.Client
}

//NewOSS
func NewSTS( accessKeyId, secret string) (*STS, error) {
	client, err := sts.NewClientWithAccessKey("cn-shanghai", accessKeyId, secret)
	if err != nil {
		return nil, err
	}
	return &STS{

		client: client,
	}, nil
}

func (s *STS) AssumeRole () (*sts.AssumeRoleResponse,error){
	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见API参考。
	request.RoleArn = "acs:ram::1855340179767769:role/osssts"
	request.RoleSessionName = helper.RandString(32)
fmt.Println(request)
	return s.client.AssumeRole(request)
}
