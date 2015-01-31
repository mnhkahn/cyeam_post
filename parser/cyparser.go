package parser

import (
	"cyeam_post/cygo"
	"cyeam_post/models"
	"fmt"
	"github.com/franela/goreq"
	"net/http"
	"strings"
	"time"
)

type CyParser struct {
	RegParser
	NormalParser
	is_debug bool
	document *CssParser
	req      *goreq.Request
}

func NewCyParser() *CyParser {
	goreq.SetConnectTimeout(time.Duration(60) * time.Second)
	p := new(CyParser)
	p.req = new(goreq.Request)
	p.req.Method = "GET"
	p.req.UserAgent = "Cyeambot"
	p.req.Timeout = time.Duration(60) * time.Second
	p.req.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
	// p.req.Compression = goreq.Gzip()
	return p
}

func (this *CyParser) Debug(is_debug bool) {
	this.is_debug = is_debug
}

func (this *CyParser) ParseHtml(post *models.Post) ([]string, error) {
	next_urls := []string{}

	var err error
	post.Source, err = this.GetHost(post.Link)
	if err != nil {
		return nil, err
	}

	this.req.Uri = post.Link
	if this.is_debug {
		this.req.ShowDebug = true
	}
	// Proxy:       "http://114.255.183.173:8080",

	res, err := this.req.Do()
	if err != nil {
		return next_urls, err
	}
	// If status code is not 200 OK, skip it
	if res.StatusCode != http.StatusOK {
		// If status code is 301, the next url is in the response Header of Location
		if res.StatusCode == http.StatusMovedPermanently {
			return []string{res.Header.Get("Location")}, nil
		} else {
			return next_urls, fmt.Errorf("Unsupported status code: %d", res.StatusCode)
		}
	}
	// If mime is not html, skip it
	if strings.Index(res.Header.Get("Content-Type"), "text/html") == -1 {
		return next_urls, fmt.Errorf("unsupported content type :%s", res.Header.Get("Content-Type"))
	}

	// 得到字符串来解析出征文
	body, err := res.Body.ToString()
	if err != nil {
		panic(err)
	}
	post.Detail = body

	this.document, err = NewCssParser(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	bchindren := this.document.GetChildren("body")
	for _, r := range bchindren {
		temp := this.RemoveHtml(r)
		if len(temp) > len(post.Description) {
			post.Description = temp
			temp_document, err := NewCssParser(strings.NewReader(r))
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
	// post.Description = this.GetMainBody(this.RemoveHtml(body))

	post.Title = this.document.GetTitle()
	next_urls = this.document.GetAs("http://" + post.Source)

	post.CreateTime = cygo.Now()
	post.ParseDate = cygo.Now()
	return next_urls, nil
}

func init() {
	Register("CyParser", NewCyParser())
}
