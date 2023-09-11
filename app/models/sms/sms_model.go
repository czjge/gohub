package sms

import "github.com/czjge/gohub/app/models"

type Sms struct {
	models.BaseModel

	Phone         string `json:"phone" gorm:"type:char(11);comment:手机号;default:'';index:'idx_phone';"`
	SignName      string `json:"sign_name" gorm:"type:varchar(255);comment:短信签名;default:'';"`
	TemplateCode  string `json:"template_code" gorm:"type:varchar(100);comment:短信模板;default:'';"`
	TemplateParam string `json:"template_param" gorm:"type:text;comment:短信模板参数;default:null;"`
	RequestId     string `json:"request_id" gorm:"type:varchar(100);comment:请求ID;default:'';uniqueIndex;"`
	BizId         string `json:"biz_id" gorm:"type:varchar(100);comment:发送回执ID;default:'';"`
	Code          string `json:"code" gorm:"type:varchar(100);comment:请求状态码;default:'';"`
	Message       string `json:"message" gorm:"type:varchar(100);comment:状态码描述;default:'';"`

	models.CommonTimestampsField
}
