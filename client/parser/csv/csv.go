package csv

import (
	"encoding/csv"
	"github.com/yaronsumel/simpleuser/server/user"
	"io"
	"os"
)

type CsvParser struct {
}

func NewCsvParser() *CsvParser {
	return &CsvParser{}
}

func (p *CsvParser) Parse(path string, uChan chan *user.Object) error {
	// open file
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// new csv reader
	r := csv.NewReader(f)
	// keep reading till error comes
	for {
		line, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if line[0] != "" && line[1] != "" {
			uChan <- &user.Object{
				Name:  line[0],
				Email: line[1],
			}
		}
	}
	return nil
}
