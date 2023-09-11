package sms

import (
	"encoding/json"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/czjge/gohub/app/models/sms"
	"github.com/czjge/gohub/pkg/logger"
)

// 需要实现 sms.Driver interface
type Aliyun struct{}

func (s *Aliyun) Send(phone string, message Mesage, config map[string]string) bool {

	client, err := CreateClient(tea.String(config["AccessKeyId"]), tea.String(config["AccessKeySecret"]))
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析绑定错误", err.Error())
		return false
	}

	params, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "短信模板参数解析错误", err.Error())
		return false
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(config["SignName"]),
		TemplateCode:  tea.String(message.Template),
		TemplateParam: tea.String(string(params)),
	}

	logger.DebugJSON("短信[阿里云]", "请求内容", sendSmsRequest)

	// 其他运行参数
	runtime := &util.RuntimeOptions{}

	_result, err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	logger.DebugJSON("短信[阿里云]", "接口响应", _result)
	if err != nil {
		var errs = &tea.SDKError{}
		if _t, ok := err.(*tea.SDKError); ok {
			errs = _t
		} else {
			errs.Message = tea.String(err.Error())
		}

		var r dysmsapi20170525.SendSmsResponseBody
		err = json.Unmarshal([]byte(*errs.Data), &r)
		logger.ErrorString("短信[阿里云]", "发送失败", err.Error())

		return false
	}

	smsModel := sms.Sms{
		Phone:         phone,
		SignName:      config["SignName"],
		TemplateCode:  message.Template,
		TemplateParam: string(params),
		RequestId:     tea.StringValue(_result.Body.RequestId),
		BizId:         tea.StringValue(_result.Body.BizId),
		Code:          tea.StringValue(_result.Body.Code),
		Message:       tea.StringValue(_result.Body.Message),
	}
	smsModel.SaveSmsLog()
	logger.DebugString("短信[阿里云]", "发信成功", "")

	return true
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}
