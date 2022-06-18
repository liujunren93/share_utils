package tencent

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type smsClient struct {
	client *sms.Client
}

func NewSmsClient(secretID, secretKey string) (cli smsClient, err error) {
	credential := common.NewCredential(secretID, secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	c, err := sms.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		return
	}
	cli.client = c
	return
}

func (c smsClient) Send(appId, signName,templateID string,mobile ,smsArgs []string) ([]byte,error){

	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(appId)
	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，签名信息可登录 [短信控制台] 查看 */
	request.SignName = common.StringPtr(signName)
	/* 用户的 session 内容: 可以携带用户侧 ID 等上下文信息，server 会原样返回 */
	//request.SessionContext = common.StringPtr("xxx")
	/* 短信码号扩展号: 默认未开通，如需开通请联系 [sms helper] */
	//request.ExtendCode = common.StringPtr("")
	/* 模板参数: 若无模板参数，则设置为空*/
	request.TemplateParamSet = common.StringPtrs(smsArgs)
	/* 模板 ID: 必须填写已审核通过的模板 ID。模板ID可登录 [短信控制台] 查看 */
	request.TemplateId = common.StringPtr(templateID)
	/* 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	 * 示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	request.PhoneNumberSet = common.StringPtrs(mobile)
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := c.client.SendSms(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return	nil,fmt.Errorf("An API error has returned: %s", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		return	nil ,err
	}
	return  json.Marshal(response.Response)

}
