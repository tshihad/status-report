package mailapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/quotedprintable"

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

//NewSender Creates new sender
func NewSender(Username, Password string) Sender {
	return Sender{Username}
}

//SendMail using simple smtp auth
func (sender Sender) sendMail(dest, cc []string, subject, msg string) error {
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
	_, err = srv.Users.Messages.Send(sender.user, message).Do()

	if err != nil {
		return errors.Wrapf(err, "Failed to send mail")
	}
	return nil
}

func (sender Sender) writeEmail(to, cc []string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = sender.user

	receipient := ""

	for _, user := range to {
		receipient = receipient + user
	}
	Cc := ""
	for _, c := range cc {
		Cc += c
	}

	header["To"] = receipient
	header["Cc"] = Cc
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

//SendHTMLEmail to write html formated mails
func (sender *Sender) SendHTMLEmail(dest, cc []string, subject, bodyMessage string) error {
	msg := setMessage(sender.user, dest, cc, subject, bodyMessage)
	return sender.sendMail(dest, cc, subject, msg)
}
