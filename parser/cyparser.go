package parser

import (
	"cyeam_post/cygo"
	"cyeam_post/models"
	"fmt"
	"github.com/franela/goreq"
	"net/http"
	"strings"
)

type CyParser struct {
	RegParser
	NormalParser
	document *CssParser
	req      *goreq.Request
}

func NewCyParser() *CyParser {
	// goreq.SetConnectTimeout(time.Duration(60) * time.Second)
	p := new(CyParser)
	p.req = new(goreq.Request)
	p.req.Method = "GET"
	p.req.UserAgent = "Cyeambot"
	// p.req.Timeout = time.Duration(60) * time.Second
	p.req.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
	// p.req.Compression = goreq.Gzip()
	return p
}

func (this *CyParser) ParseHtml(post *models.Post) ([]string, error) {
	next_urls := []string{}

	var err error
	post.Source, err = this.GetHost(post.Link)
	if err != nil {
		return nil, err
	}

	this.req.Uri = post.Link

	// ShowDebug: true,
	// Proxy:       "http://114.255.183.173:8080",

	res, err := this.req.Do()
	if err != nil || res.StatusCode != http.StatusOK {
		return next_urls, fmt.Errorf("Unsupported status code: %d", res.StatusCode)
	}

	// 得到字符串来解析出征文
	body, err := res.Body.ToString()
	if err != nil {
		panic(err)
	}
	post.Detail = body
	post.Description = this.GetMainBody(this.RemoveHtml(body))
	this.document, err = NewCssParser(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	post.Title = this.document.GetTitle()
	imgs := this.document.GetImgs()
	if len(imgs) > 0 {
		if strings.HasPrefix(imgs[0], "/") {
			post.Figure = "http://" + post.Source + imgs[0]
		} else {
			post.Figure = imgs[0]
		}
	}
	next_urls = this.document.GetAs("http://" + post.Source)

	post.CreateTime = cygo.Now()
	post.ParseDate = cygo.Now()
	return next_urls, nil
}

func init() {
	Register("CyParser", NewCyParser())
}
