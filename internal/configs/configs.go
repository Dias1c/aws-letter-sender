package configs

import (
	"fmt"
	"runtime"
)

const ProgramName = "aws-ses-bulk-emails"
const SourceCodeLink = "https://github.com/Dias1c/aws-ses-bulk-emails"

var Vesrion = "x.y.z_dev"

func GetExpectedProgramFullName() string {
	return fmt.Sprintf("%v-v%v-%v-%v", ProgramName, Vesrion, runtime.GOOS, runtime.GOARCH)
}
