package parser

import (
	"net/url"
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
