package sms

import "github.com/czjge/gohub/app/models"

type Sms struct {
	models.BaseModel

	Phone         string `json:"phone"`
	SignName      string `json:"sign_name"`
	TemplateCode  string `json:"template_code"`
	TemplateParam string `json:"template_param"`
	RequestId     string `json:"request_id"`
	BizId         string `json:"biz_id"`
	Code          string `json:"code"`
	Message       string `json:"message"`

	models.CommonTimestampsField
}
