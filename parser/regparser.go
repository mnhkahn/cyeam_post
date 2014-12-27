package parser

import (
	"regexp"
)

type RegParser struct {
}

func (this *RegParser) RemoveHtml(src string) string {
	re_html := regexp.MustCompile("(?is)<.*?>")
	return re_html.ReplaceAllString(src, "")
}
