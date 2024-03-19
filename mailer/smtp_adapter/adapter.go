package smtp_adapter

import (
	"bytes"
	"context"
	"crypto/tls"
	"go-monolith-template/mailer"
	"go-monolith-template/templates"
	gomail "gopkg.in/mail.v2"
)

type Options struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func New(opts Options) mailer.Emailer {
	m := SMTPMailer{
		host:     opts.Host,
		port:     opts.Port,
		username: opts.Username,
		password: opts.Password,
		from:     opts.From,
	}
	return m
}

type SMTPMailer struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func (s SMTPMailer) SendEmail(ctx context.Context, input mailer.SendEmailInput) error {
	d := gomail.NewDialer(s.host, s.port, s.username, s.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", input.To)
	m.SetHeader("Subject", input.Subject)
	var b bytes.Buffer
	templates.BaseEmail(templates.BaseEmailOptions{
		Title:            input.Title,
		PoweredByLink:    input.PoweredByLink,
		PoweredByText:    input.PoweredByText,
		ContentHeader:    input.ContentHeader,
		ContentText:      input.ContentText,
		CallToActionLink: input.CallToActionLink,
		CallToActionText: input.CallToActionText,
		UnsubscribeLink:  input.UnsubscribeLink,
	}).Render(ctx, &b)
	content := b.String()
	m.SetBody("text/html", content)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
