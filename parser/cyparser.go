package parser

import (
	"cyeam_post/cygo"
	"cyeam_post/models"
	"fmt"
	"github.com/franela/goreq"
	"time"
)

type CyParser struct {
	RegParser
	NormalParser
}

func (this *CyParser) ParseHtml(post *models.Post) ([]string, error) {
	var err error
	post.Source, err = this.GetHost(post.Link)
	if err != nil {
		return nil, err
	}
	// post.Link = this.GetUrl(post.Link)
	req := goreq.Request{
		Method:      "GET",
		Uri:         post.Link,
		ContentType: "application/json",
		UserAgent:   "Cyeambot",
		Timeout:     time.Duration(60) * time.Second,
		// Proxy:       "http://114.255.183.173:8080",
		Compression: goreq.Gzip(),
	}
	req.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
	goreq.SetConnectTimeout(time.Duration(60) * time.Second)
	res, err := req.Do()
	if err != nil {
		panic(err)
	}

	body, err := res.Body.ToString()
	fmt.Println(body, err)
	if err != nil {
		return nil, err
	}
	// post.Title =
	post.CreateTime = cygo.Now()
	post.Author = "Cyeam"
	post.Detail = body
	imgs := this.GetImgs(body)
	// fmt.Println(imgs)
	if len(imgs) > 0 {
		post.Figure = imgs[0]
	}
	// post.Description =
	post.ParseDate = cygo.Now()
	next_urls := this.GetAs(body, "http://"+post.Source)
	return next_urls, nil
}

func init() {
	Register("CyParser", &CyParser{})
}
