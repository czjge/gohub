package mail

import (
	"sync"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/helpers"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte
	HTML    []byte
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})

	return internalMailer
}

func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email, helpers.Struct2Map(config.GetConfig().Email.Smtp))
}
