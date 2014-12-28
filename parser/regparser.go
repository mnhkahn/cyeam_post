package parser

import (
	"regexp"
	"strings"
)

type RegParser struct {
}

var (
	RE_HOST = regexp.MustCompile(`http(s)?://([\w-]+\.)+[\w-]+/?`)
	RE_HTML = regexp.MustCompile("(?is)<.*?>")
	RE_A    = regexp.MustCompile("(?is)<a .*?</a>")
	RE_HREF = regexp.MustCompile(`href=\"?(.*?)(\"|>|\\s+)`)
	RE_IMG  = regexp.MustCompile("(?is)<img .*?>")
	RE_SRC  = regexp.MustCompile(`src=\"?(.*?)(\"|>|\\s+)`)
)

func (this *RegParser) RemoveHtml(src string) string {
	return RE_HTML.ReplaceAllString(src, "")
}

func (this *RegParser) GetAs(body string) []string {
	next_urls := RE_A.FindAllString(body, -1)
	for i := 0; i < len(next_urls); i++ {
		temp := RE_HREF.Find([]byte(next_urls[i]))
		if len(temp) > 0 {
			temp = temp[strings.Index(string(temp), "href")+6 : len(temp)-1]
			next_urls[i] = string(temp)
		}
	}
	return next_urls
}

func (this *RegParser) GetImgs(body string) []string {
	return RE_IMG.FindAllString(body, -1)
}
