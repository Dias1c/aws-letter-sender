package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/Dias1c/aws-letter-sender/pkg/desktop"
)

func main() {
	err := desktop.Run()
	switch {
	case err == nil:
	case errors.Is(err, desktop.ErrFlagsRequired):
		log.Println(err)
		flag.PrintDefaults()
		os.Exit(1)
	case err != nil:
		log.Fatal(err)
	}
}
