package sender

import (
	"os"

	"github.com/subosito/gotenv"
)

const (
	ENV_AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
	ENV_AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	ENV_AWS_SENDER_EMAIL      = "AWS_SENDER_EMAIL"
	ENV_AWS_SENDER_REGION     = "AWS_SENDER_REGION"
)

type Config struct {
	SecretAccessKey string
	AccessKeyID     string
	SenderEmail     string
	SenderRegion    string
}

func GetAWSParams() *Config {
	gotenv.Load()
	conf := &Config{
		SecretAccessKey: os.Getenv(ENV_AWS_SECRET_ACCESS_KEY),
		AccessKeyID:     os.Getenv(ENV_AWS_ACCESS_KEY_ID),
		SenderEmail:     os.Getenv(ENV_AWS_SENDER_EMAIL),
		SenderRegion:    os.Getenv(ENV_AWS_SENDER_REGION),
	}
	return conf
}
