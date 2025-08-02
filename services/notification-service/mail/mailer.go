package mail

import (
	"log"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	To       string
}

func (m *Mailer) Send(subject, body string) {
	message := gomail.NewMessage()
	message.SetHeader("From", m.From)
	message.SetHeader("To", m.To)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)
	if err := dialer.DialAndSend(message); err != nil {
		log.Println("Failed to send email:", err)
	} else {
		log.Println("Email sent successfully!")
	}
}
