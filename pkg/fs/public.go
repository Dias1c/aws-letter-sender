package fs

import "strings"

func GetExtension(filepath string) string {
	splitted := strings.Split(filepath, ".")
	return splitted[len(splitted)-1]
}
