package desktop

import (
	"fmt"

	"github.com/Dias1c/aws-ses-bulk-emails/internal/configs"
)

func printProgramVersion() {
	fmt.Printf("%s\n", configs.GetExpectedProgramFullName())
}
