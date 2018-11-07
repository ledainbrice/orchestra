package config

import (
	"log"

	"github.com/mailgun/mailgun-go"
)

type Reciever struct {
	email      string
	parameters map[string]interface{}
}

func NewReciever(email string, params map[string]interface{}) Reciever {
	var reciever = Reciever{
		email:      email,
		parameters: params,
	}
	return reciever
}
func SendEmail(key string, subject string, body string, mailingList []Reciever) error {
	conf := env.emails[key]
	mg := mailgun.NewMailgun(conf.domain, conf.apiKey)
	m := mg.NewMessage(
		conf.mail,
		subject,
		body,
	)
	for _, reciever := range mailingList {
		err := m.AddRecipientAndVariables(reciever.email, reciever.parameters)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, _, err := mg.Send(m)

	return err
}
