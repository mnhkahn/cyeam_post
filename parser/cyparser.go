package parser

import (
	// . "cyeam_post/logs"
	"cyeam_post/models"
	"github.com/astaxie/beego/httplib"
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
	req := httplib.Get(post.Link)
	req.SetUserAgent("Cyeambot")
	req.SetTimeout(time.Second*5, time.Second*5)
	body, err := req.String()
	if err != nil {
		return nil, err
	}
	post.Detail = body
	imgs := this.GetImgs(body)
	// fmt.Println(imgs)
	if len(imgs) > 0 {
		post.Figure = imgs[0]
	}
	next_urls := this.GetAs(body, "http://"+post.Source)
	return next_urls, nil
}

func init() {
	Register("CyParser", &CyParser{})
}
