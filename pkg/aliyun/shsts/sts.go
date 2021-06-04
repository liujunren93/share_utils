package shsts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/liujunren93/share_utils/helper"
)

type STS struct {
	client *sts.Client
}

//NewOSS
func NewSTS(accessKeyId, secret string) (*STS, error) {
	client, err := sts.NewClientWithAccessKey("cn-shanghai", accessKeyId, secret)

	if err != nil {
		return nil, err
	}
	return &STS{
		client: client,
	}, nil
}

//Credentials 获取访问令牌
//@return Credentials 令牌
//@return string sessionName
//@return error
func (s *STS) Credentials(roleArn, session string) (*sts.Credentials, string, error) {
	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见API参考。

	request.RoleArn = roleArn
	if session == "" {
		session = helper.RandString(32)
	}
	request.RoleSessionName = session

	role, err := s.client.AssumeRole(request)
	if err != nil {
		return nil, "", err
	}
	return &role.Credentials,session, nil

}
