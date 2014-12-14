package parser

import (
	"cyeam_post/models"
	"fmt"
	"regexp"
)

type ParserContainer interface {
	Index(i int) *models.Post
	Set(i int, p *models.Post) *models.Post
	Len() int
}

type Parser interface {
	Parse(source string) (ParserContainer, error)
}

var parsers = make(map[string]Parser)

// Register makes a config adapter available by the adapter name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, parser Parser) {
	if parser == nil {
		panic("parser: Register parser is nil")
	}
	if _, ok := parsers[name]; ok {
		panic("parser: Register called twice for adapter " + name)
	}
	parsers[name] = parser
}

func NewParser(parser_name, source string) (ParserContainer, error) {
	parser, ok := parsers[parser_name]
	if !ok {
		return nil, fmt.Errorf("parser: unknown parser_name %q", parser_name)
	}
	return parser.Parse(source)
}

func RemoveHtml(src string) string {
	re_html := regexp.MustCompile("(?is)<.*?>")
	return re_html.ReplaceAllString(src, "")
}
