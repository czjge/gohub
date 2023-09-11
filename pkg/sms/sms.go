package sms

import (
	"sync"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/helpers"
)

// 短信结构体
type Mesage struct {
	Template string
	Data     map[string]string
	Content  string
}

// 发送短信操作类
type SMS struct {
	Driver Driver
}

// 单例模式
var once sync.Once

var internalSMS *SMS

func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})

	return internalSMS
}

func (sms *SMS) Send(phone string, message Mesage) bool {
	return sms.Driver.Send(phone, message, helpers.Struct2Map(config.GetConfig().Sms))
}
