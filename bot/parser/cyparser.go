package parser

import (
	"cyeam_post/bot/parser/selector"
	"cyeam_post/cygo"
	"cyeam_post/models"
	"strings"
)

type CyParser struct {
	NormalParser
	css *selector.CssSelector
	reg *selector.RegSelector
}

func (this *CyParser) ParseHtml(post *models.Post, body string) ([]string, error) {
	next_urls := []string{}
	var err error

	post.Detail = body
	post.Source, err = this.GetHost(post.Link)
	if err != nil {
		return nil, err
	}

	this.css, err = selector.NewCssSelector(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	bchindren := this.css.GetChildren("body")
	for _, r := range bchindren {
		temp := this.reg.RemoveHtml(r)
		if len(temp) > len(post.Description) {
			post.Description = temp
			temp_document, err := selector.NewCssSelector(strings.NewReader(r))
			if err != nil {
				return nil, err
			}
			imgs := temp_document.GetImgs()
			if len(imgs) > 0 {
				if strings.HasPrefix(imgs[0], "/") {
					post.Figure = "http://" + post.Source + imgs[0]
				} else {
					post.Figure = imgs[0]
				}
			}
		}
	}
	post.Title = this.css.GetTitle()
	next_urls = this.css.GetAs("http://" + post.Source)

	post.CreateTime = cygo.Now()
	post.ParseDate = cygo.Now()
	return next_urls, nil
}

func init() {
	Register("CyParser", &CyParser{})
}
