package desktop

import (
	"flag"
	"fmt"

	"github.com/Dias1c/aws-ses-bulk-emails/internal/configs"
)

func printProgramUsage() {
	fmt.Println("Program Name:", configs.GetExpectedProgramFullName())
	fmt.Println("Usage:", configs.SourceCodeLink)
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}
