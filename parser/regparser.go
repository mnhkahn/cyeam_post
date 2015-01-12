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
	RE_A    = regexp.MustCompile(`<a(.*?)href="(.*?)"(.*?)>(.*?)</a>`)
	RE_HREF = regexp.MustCompile(`href=\"?(.*?)(\"|>|\\s+)`)
	RE_IMG  = regexp.MustCompile("(?is)<img .*?>")
	RE_SRC  = regexp.MustCompile(`src=\"?(.*?)(\"|>|\\s+)`)
)

func (this *RegParser) RemoveHtml(src string) string {
	return RE_HTML.ReplaceAllString(src, "")
}

func (this *RegParser) GetAs(body, host string) []string {
	next_urls := RE_A.FindAllString(body, -1)
	for i := 0; i < len(next_urls); i++ {
		temp := RE_HREF.Find([]byte(next_urls[i]))
		if len(temp) > 0 {
			temp = temp[strings.Index(string(temp), "href")+6 : len(temp)-1]
			next_urls[i] = string(temp)
			// 如果url以/开头，需要拼接上http协议头和域名
			if next_urls[i] == "" || strings.HasPrefix(next_urls[i], "/") {
				next_urls[i] = host + next_urls[i]
			}
			// 如果url中包含标签#，需要将标签删掉
			next_urls[i] = strings.Split(next_urls[i], "#")[0]
			// 如果url没有http头，加上
			if !strings.HasPrefix(next_urls[i], "http") {
				next_urls[i] = "http://" + next_urls[i]
			}
		}
	}
	return next_urls
}

func (this *RegParser) GetImgs(body string) []string {
	return RE_IMG.FindAllString(body, -1)
}
