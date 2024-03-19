package mailer

import "context"

type SendEmailInput struct {
	To               string
	Subject          string
	CC               string
	BCC              string
	Title            string
	PoweredByLink    string
	PoweredByText    string
	ContentHeader    string
	ContentText      string
	CallToActionLink string
	CallToActionText string
	UnsubscribeLink  string
}

type Emailer interface {
	SendEmail(ctx context.Context, input SendEmailInput) error
}
