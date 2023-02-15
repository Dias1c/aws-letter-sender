package sender

import (
	"fmt"

	"github.com/Dias1c/aws-letter-sender/pkg/letter/sender/html"
	"github.com/Dias1c/aws-letter-sender/pkg/letter/sender/text"
)

type IAWSSender interface {
	SendLetter(email, subject, body string) (string, error)
	SendLetterTemplate(email, subject, tmplpath string, data interface{}) (string, error)
}

func NewSender(email, region, format string) (IAWSSender, error) {
	switch {
	case format == ".txt" || format == ".text":
		return text.NewSender(email, region), nil
	case format == ".html":
		return html.NewSender(email, region), nil
	default:
		return nil, fmt.Errorf("%w: '%v'", ErrUndefinedSender, format)
	}
}
