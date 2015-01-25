package parser

import (
	"bufio"
	"fmt"
	"net/url"
	"strings"
)

type NormalParser struct {
}

func (this *NormalParser) GetUrl(src string) string {
	return src
}
func (this *NormalParser) GetHost(src string) (string, error) {
	u, err := url.Parse(src)
	return u.Host, err
}

func (this *NormalParser) GetMainBody(body string) (string, error) {
	result := ""
	scanner := bufio.NewScanner(strings.NewReader(body))

	_limitCount := 180
	_limitGap := 10

	gap := 0
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			gap++
		}
		if gap > _limitGap && len(result) > 0 {
			break
		}
		// 提取正文
		if len(text) > _limitCount {
			if len(result) > 0 {
				result += fmt.Sprintf("\n")
			}
			result += text
			gap = 0
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return result, nil
}
