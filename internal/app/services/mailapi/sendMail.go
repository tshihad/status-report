package mailapi

import (
	"io/ioutil"

	"google.golang.org/api/gmail/v1"

	"golang.org/x/oauth2"

	"github.com/pkg/errors"
	"golang.org/x/oauth2/google"
)

func authenticate(scope ...string) (*oauth2.Config, error) {
	b, err := ioutil.ReadFile("config/credentials.json")
	if err != nil {
		return nil, errors.Wrap(err, "Cant locate credentials. Please download credentials from google gmail api")
	}
	config, err := google.ConfigFromJSON(b, scope...)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse client secret file to config")
	}
	return config, nil
}

//NewMail Creates new sender
func NewMail(Username string) *Mail {
	return &Mail{Username}
}

//SendMail using simple smtp auth
func (mail *Mail) sendMail(dest, cc []string, subject, msg string) error {
	config, err := authenticate(gmail.GmailSendScope)
	if err != nil {
		return err
	}

	client := getClient(config)
	srv, err := gmail.New(client)
	if err != nil {
		return errors.Wrap(err, "Unable to retrieve Gmail client")
	}
	message := &gmail.Message{
		Raw: msg,
	}
	_, err = srv.Users.Messages.Send(mail.user, message).Do()

	if err != nil {
		return errors.Wrapf(err, "Failed to send mail")
	}
	return nil
}

//SendHTMLEmail to write html formated mails
func (mail *Mail) SendHTMLEmail(dest, cc []string, subject, bodyMessage string) error {
	msg := setMessage(mail.user, dest, cc, subject, bodyMessage)
	return mail.sendMail(dest, cc, subject, msg)
}
