package csv

import (
	"encoding/csv"
	"fmt"
	"io"
)

type DataCompiler struct {
}

func NewDataCompiler() *DataCompiler {
	return &DataCompiler{}
}

func (c *DataCompiler) CompileData(r io.Reader) ([]map[string]string, error) {
	csvReader := csv.NewReader(r)

	fields, err := csvReader.Read()
	if err == io.EOF {
		return nil, fmt.Errorf("no keys")
	} else if err != nil {
		return nil, fmt.Errorf("r.Read: %w", err)
	}
	var keys []string = fields

	var result []map[string]string
	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("r.Read: %w", err)
		}

		mp := make(map[string]string, len(keys))
		for i, v := range keys {
			mp[v] = fields[i]
		}

		result = append(result, mp)
	}

	return result, nil
}
