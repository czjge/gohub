package sms

type Driver interface {
	Send(phone string, message Mesage, config map[string]string) bool
}
