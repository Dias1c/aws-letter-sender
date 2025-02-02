package configs

import (
	"fmt"
	"runtime"
)

var Vesrion = "x.y.z_dev"
var VersionSourceCodeLink = fmt.Sprintf("https://github.com/Dias1c/aws-ses-bulk-emails/releases/tag/v%s", Vesrion)

const SourceCodeLink = "https://github.com/Dias1c/aws-ses-bulk-emails"
const ProgramName = "aws-ses-bulk-emails"

func GetExpectedProgramFullName() string {
	return fmt.Sprintf("%v-v%v-%v-%v", ProgramName, Vesrion, runtime.GOOS, runtime.GOARCH)
}
