package utils

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

type CssSelector struct {
	Document *goquery.Document
}

func NewCssSelector(r io.Reader) (*CssSelector, error) {
	parser := new(CssSelector)
	var err error
	parser.Document, err = goquery.NewDocumentFromReader(r)
	return parser, err
}

// h1就是文章标题，如果没有搜到，用title标签
func (this *CssSelector) GetTitle() string {
	if this.Document.Find("h1").Length() == 1 {
		return this.Document.Find("h1").Text()
	} else {
		return this.Document.Find("title").Text()
	}
}

func (this *CssSelector) GetImgs() []string {
	imgs := []string{}
	this.Document.Find("img").Each(func(i int, s *goquery.Selection) {
		src, has := s.Attr("src")
		if has {
			imgs = append(imgs, src)
		}
	})
	return imgs
}

func (this *CssSelector) GetAs(host string) []string {
	as := []string{}
	this.Document.Find("a").Each(func(i int, s *goquery.Selection) {
		href, has := s.Attr("href")
		if has {
			// 如果url以/开头，需要拼接上http协议头和域名
			if href == "" || strings.HasPrefix(href, "/") {
				href = host + href
			}
			// 如果url中包含标签#，需要将标签删掉
			href = strings.Split(href, "#")[0]
			// 如果url没有http头，加上
			if !strings.HasPrefix(href, "http") {
				href = "http://" + href
			}
			as = append(as, href)
		}
	})
	return as
}

func (this *CssSelector) GetChildren(pattern string) []string {
	res := []string{}
	this.Document.Find(pattern).Children().Each(func(i int, s *goquery.Selection) {
		html, err := s.Html()
		if err == nil {
			res = append(res, html)
		}
	})
	return res
}
