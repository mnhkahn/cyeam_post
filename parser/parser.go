package parser

import (
	"cyeam_post/models"
	"fmt"
)

type Parser interface {
	ParseHtml(post *models.Post) ([]string, error)
}

var parsers = make(map[string]Parser)

func Register(name string, parser Parser) {
	if parser == nil {
		panic("parser: Register parser is nil")
	}
	if _, ok := parsers[name]; ok {
		panic("parser: Register called twice for adapter " + name)
	}
	parsers[name] = parser
}

func NewParser(parser_name string) (Parser, error) {
	parser, ok := parsers[parser_name]
	if !ok {
		return nil, fmt.Errorf("parser: unknown parser_name %q", parser_name)
	}
	return parser, nil
}
