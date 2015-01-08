package parser

import (
// "regexp"
)

type NormalParser struct {
}

func (this *NormalParser) GetUrl(src string) string {
	return src
}
