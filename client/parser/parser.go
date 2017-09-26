package parser

import (
	"github.com/yaronsumel/simpleuser/client/parser/csv"
	"github.com/yaronsumel/simpleuser/server/user"
)

type Parser interface {
	Parse(path string, uChan chan *user.Object) error
}

// NewParser return new csv parser as default
func NewParser() Parser {
	return csv.NewCsvParser()
}
