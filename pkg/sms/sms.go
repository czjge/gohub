package sms

// 短信结构体
type Mesage struct {
	Template string
	Data     map[string]string
	Content  string
}
