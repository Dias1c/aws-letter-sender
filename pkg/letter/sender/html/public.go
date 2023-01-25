package html

import (
	"bytes"
	"fmt"

	tmpl "html/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Sender struct {
	email  string // example@example.com
	region string // eu-west-1
}

func NewSender(awsemail, awsregion string) *Sender {
	return &Sender{
		email:  awsemail,
		region: awsregion,
	}
}

func (s *Sender) SendLetterTemplate(email, subject, tmplpath string, data interface{}) (string, error) {
	t, err := tmpl.ParseFiles(tmplpath)
	if err != nil {
		return "", fmt.Errorf("ParseFiles: %v", err)
	}

	var body bytes.Buffer
	err = t.Execute(&body, data)
	if err != nil {
		return "", fmt.Errorf("Execute: %v", err)
	}

	msgID, err := s.SendLetter(email, subject, body.String())
	if err != nil {
		return "", fmt.Errorf("SendLetter: %v", err)
	}
	return msgID, err
}

func (s *Sender) SendLetter(email, subject, htmlbody string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s.region)},
	)
	if err != nil {
		return "", fmt.Errorf("session.NewSession: %s", err)
	}

	var charset string = "UTF-8"
	// Create an SES session.
	svc := ses.New(sess)
	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charset),
					Data:    aws.String(htmlbody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charset),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(s.email),
	}

	// Attempt to send the email.
	sent, err := svc.SendEmail(input)
	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				return "", fmt.Errorf("svc.SendEmail: " + ses.ErrCodeMessageRejected + aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				return "", fmt.Errorf("svc.SendEmail: "+ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				return "", fmt.Errorf("svc.SendEmail: "+ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				return "", fmt.Errorf("svc.SendEmail: " + aerr.Error())
			}
		} else {
			return "", fmt.Errorf("svc.SendEmail: " + err.Error())
		}
	}

	return *sent.MessageId, nil
}
