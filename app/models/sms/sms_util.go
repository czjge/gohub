package sms

import "github.com/czjge/gohub/pkg/database"

func (sms *Sms) SaveSmsLog() {
	database.DB().Create(&sms)
}
