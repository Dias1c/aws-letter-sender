package desktop

import "fmt"

var (
	ErrFlagsRequired = fmt.Errorf("flags requred")
	ErrShowVersion   = fmt.Errorf("show version")
)
