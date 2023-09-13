package sendgrid

import (
	"web3Server/config"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var mailer *sendgrid.Client

func Start() {
	sendgrid_api_key := config.Config.GetString("sendgrid.api_key")
	mailer = sendgrid.NewSendClient(sendgrid_api_key)
}

func SendLoginEmail(toemail, code string) error {
	fromemail := config.Config.GetString("sendgrid.from")
	from := mail.NewEmail(fromemail, fromemail)
	subject := "Welcome to Guan Dan!"
	to := mail.NewEmail("Hey Player", toemail)
	plainTextContent := "Welcome to Guan Dan! To login, you’ll need to use your email and verify code:" + code
	htmlContent := ""
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	_, err := mailer.Send(message)
	return err
}
