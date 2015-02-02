package parser

import (
	"cyeam_post/models"
	"fmt"
	"net/url"
	"strings"
)

type NormalParser struct {
	is_debug bool
}

func (this *NormalParser) ParseHtml(post *models.Post) ([]string, error) {
	return nil, nil
}

func (this *NormalParser) Debug(is_debug bool) {
	this.is_debug = is_debug
}

func (this *NormalParser) GetUrl(src string) string {
	return src
}
func (this *NormalParser) GetHost(src string) (string, error) {
	u, err := url.Parse(src)
	return u.Host, err
}

func (this *NormalParser) GetMainBody(body string) string {
	result := ""
	_limitCount := 180
	_limitGap := 10

	texts := strings.Split(body, "\n")
	gap := 0
	for _, text := range texts {
		if len(text) == 0 {
			gap++
		}
		if gap > _limitGap && len(result) > 0 {
			break
		}
		// 提取正文
		if len(text) > _limitCount || strings.Index(text, "<code") != -1 {
			if len(result) > 0 {
				result += fmt.Sprintf("\n")
			}
			result += text
			gap = 0
		}
	}

	return strings.Trim(result, " ")
}