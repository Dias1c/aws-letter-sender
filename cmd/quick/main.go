package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Dias1c/aws-letter-sender/internal/letter/dtcomp"
	"github.com/Dias1c/aws-letter-sender/internal/letter/sender"
	"github.com/Dias1c/aws-letter-sender/pkg/fs"
)

type Params struct {
	DataExt  string
	DataFile string

	EmailSender string
	EmailRegion string

	LetterSubject string

	TemplateFile string
	TemplateExt  string
}

func main() {
	params, err := GetParams()
	if err != nil {
		flag.PrintDefaults()
		log.Fatalf("GetParams: %v", err)
	}

	comp, err := dtcomp.GetDataCompiler(params.DataExt)
	if err != nil {
		log.Fatalf("dtcomp.GetDataCompiler: %v", err)
	}

	f, err := os.OpenFile(params.DataFile, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatalf("os.Openfile: %v", err)
	}
	lines, err := comp.CompileData(f)
	if err != nil {
		log.Fatalf("comp.CompileData: %v", err)
	}

	lsender, err := sender.NewSender(params.EmailSender, params.EmailRegion, params.TemplateExt)
	if err != nil {
		log.Fatalf("sender.NewSender: %v", err)
	}
	for i := 0; i < len(lines); i++ {
		mp := lines[i]
		email := mp["email"]
		subject := params.LetterSubject
		if value, ok := mp["subject"]; ok {
			subject = value
		}
		template := params.TemplateFile
		if value, ok := mp["template"]; ok {
			template = value
		}

		res, err := lsender.SendLetterTemplate(email, subject, template, mp)
		fmt.Println(res, err)
	}

}

func GetParams() (*Params, error) {
	// Priority of params from
	// 1. Data file
	// 2. Args
	// 3. .env file
	awsParams := sender.GetAWSParams()
	params := &Params{
		EmailSender: awsParams.SenderEmail,
		EmailRegion: awsParams.SenderRegion,
	}

	emailSender := flag.String("email-sender", "", "")                                                    // --email-sender=""
	subject := flag.String("subject", "", `letter subject value, if empty, takes subject from data file`) // --subject="name" - subject
	dataExt := flag.String("data-ext", "", `optional, if empty sets mode auto`)                           // --data-ext="csv" - Will use extencion which set
	dataFile := flag.String("data-file", "", `required, `)                                                // --data-file="filepath" - path of data ext file
	tmplFile := flag.String("tmpl-file", "", ``)                                                          // --tmpl-file="path" -
	tmplExt := flag.String("tmpl-ext", "", `if empty sets mode auto`)
	flag.Parse()

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

	if *dataExt != "" {
		params.DataExt = *dataExt
	} else {
		params.DataExt = fs.GetExtension(params.DataFile)
	}

	if *tmplFile != "" {
		params.TemplateFile = *tmplFile
	}

	if *tmplExt != "" {
		params.TemplateExt = *tmplExt
	} else {
		params.TemplateExt = fs.GetExtension(params.TemplateFile)
	}
	return params, nil
}
