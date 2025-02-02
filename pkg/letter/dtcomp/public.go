package dtcomp

import (
	"fmt"
	"io"

	"github.com/Dias1c/aws-ses-bulk-emails/pkg/letter/dtcomp/csv"
)

type IDataCompiler interface {
	CompileData(r io.Reader) ([]map[string]string, error)
}

func GetDataCompiler(extension string) (IDataCompiler, error) {
	switch {
	case extension == ".csv":
		return csv.NewDataCompiler(), nil
	default:
		return nil, fmt.Errorf("struct for extension '%v' not found", extension)
	}
}
