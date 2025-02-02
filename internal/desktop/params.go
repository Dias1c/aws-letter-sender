package desktop

import (
	"errors"
	"flag"
	"fmt"
	"runtime"

	"github.com/Dias1c/aws-ses-bulk-emails/internal/configs"
	"github.com/Dias1c/aws-ses-bulk-emails/pkg/letter/sender"
)

type Params struct {
	DataFile string

	EmailSender string
	EmailRegion string

	LetterSubject string

	TemplateFile string
}

func GetParams() (*Params, error) {
	awsParams := sender.GetAWSParams()
	params := &Params{
		EmailSender: awsParams.SenderEmail,
		EmailRegion: awsParams.SenderRegion,
	}

	var (
		showVersion = flag.Bool("version", false, "Show version")

		emailSender = flag.String("email-sender", "", `[optional]
aws sender email
`)  // --email-sender="
		subject = flag.String("subject", "", `[optional] 
letter subject value, if empty, takes subject 
from data file
`)  // --subject="name" - subject
		dataFile = flag.String("data-file", "data.csv", fmt.Sprintf(`[optional]
file which stores emails to send and data for 
using template. It must contain values for key
'EMAIL'. If file contain 
%q keys,
it will use them as main params to send letter.

By default equal to "data.csv"
`, lineVars))  // --data-file="filepath" - path of data ext file
		tmplFile = flag.String("tmpl-file", "", `[optional]
template file. Files must end with '.txt', 
'.text', '.html'
`)  // --tmpl-file="path" -
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("%v-v%v-%v-%v\n", configs.ProgramName, configs.Vesrion, runtime.GOOS, runtime.GOARCH)
		return nil, ErrShowVersion
	}

	if *emailSender != "" {
		params.EmailSender = *emailSender
	}
	if *subject != "" {
		params.LetterSubject = *subject
	}

	if *dataFile == "" {
		return nil, errors.New("no data file")
	} else {
		params.DataFile = *dataFile
	}

	if *tmplFile != "" {
		params.TemplateFile = *tmplFile
	}

	return params, nil
}
