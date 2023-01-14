package desktop

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Dias1c/aws-letter-sender/internal/letter/dtcomp"
	"github.com/Dias1c/aws-letter-sender/internal/letter/sender"
	"github.com/Dias1c/aws-letter-sender/pkg/fs"
)

func init() {
	logFile, err := os.OpenFile(".history.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	SentLogger = log.New(mw, "", log.LstdFlags)
}

const (
	LV_EMAIL         = "EMAIL"
	LV_SUBJECT       = "SUBJECT"
	LV_TEMPLATE_FILE = "TEMPLATE_FILE"
	LV_SENDER_EMAIL  = "SENDER_EMAIL"
	LV_SENDER_REGION = "SENDER_REGION"
)

var SentLogger *log.Logger
var lineVars = []string{
	LV_EMAIL,
	LV_SUBJECT,
	LV_TEMPLATE_FILE,
	LV_SENDER_EMAIL,
	LV_SENDER_REGION,
}

// Run - runs program
// Priority of params
// 1. Data file
// 2. Args
// 3. .env file
func Run() error {
	params, err := GetParams()
	if err != nil {
		return fmt.Errorf("GetParams: %w(%v)", ErrFlagsRequired, err)
	}

	comp, err := dtcomp.GetDataCompiler(fs.GetExtension(params.DataFile))
	if err != nil {
		return fmt.Errorf("dtcomp.GetDataCompiler: %w", err)
	}

	f, err := os.OpenFile(params.DataFile, os.O_RDONLY, 0777)
	if err != nil {
		return fmt.Errorf("os.Openfile: %w", err)
	}

	lines, err := comp.CompileData(f)
	if err != nil {
		return fmt.Errorf("comp.CompileData: %w", err)
	}

	for i := 0; i < len(lines); i++ {
		msg, err := SendForLine(lines[i], params)
		if err != nil {
			return fmt.Errorf("SendForLine: %w", err)
		}
		SentLogger.Println(msg)
	}
	return nil
}

func SendForLine(mp map[string]string, params *Params) (msg string, err error) {
	email := mp[LV_EMAIL]
	subject := params.LetterSubject
	if value, ok := mp[LV_SUBJECT]; ok {
		subject = value
	}
	template := params.TemplateFile
	if value, ok := mp[LV_TEMPLATE_FILE]; ok {
		template = value
	}
	senderEmail := params.EmailSender
	if value, ok := mp[LV_SENDER_EMAIL]; ok {
		senderEmail = value
	}
	senderRegion := params.EmailRegion
	if value, ok := mp[LV_SENDER_REGION]; ok {
		senderRegion = value
	}
	templateExt := fs.GetExtension(template)

	lsender, err := sender.NewSender(senderEmail, senderRegion, templateExt)
	if err != nil {
		return "", fmt.Errorf("sender.NewSender: %w", err)
	}

	res, err := lsender.SendLetterTemplate(email, subject, template, mp)
	msg = fmt.Sprintf("%v %q subject:%q template:%q err:\"%v\"", res, email, subject, template, err)
	return msg, nil
}
