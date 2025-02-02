package desktop

import (
	"flag"
	"fmt"

	"github.com/Dias1c/aws-ses-bulk-emails/internal/configs"
)

func printProgramUsage() {
	fmt.Println("Program Name:", configs.GetExpectedProgramFullName())
	fmt.Println("Description:  Bulk emails using AWS SES with golang HTML/text templates support")
	fmt.Println("Release:      ", configs.VersionSourceCodeLink)
	fmt.Println("Usage:        ", configs.SourceCodeLink)
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}
