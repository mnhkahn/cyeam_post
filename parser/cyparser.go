package parser

import (
	"cyeam_post/models"
	// "fmt"
	. "cyeam_post/logs"
	"github.com/astaxie/beego/httplib"
)

type CyParser struct {
	RegParser
	NormalParser
}

func (this *CyParser) ParseHtml(post *models.Post) ([]string, error) {
	Log.Info(post.Link)
	// post.Link = this.GetUrl(post.Link)
	req := httplib.Get(post.Link)
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
	next_urls := this.GetAs(body)
	return next_urls, nil
}

func init() {
	Register("CyParser", &CyParser{})
}
