package parser

import (
	"cyeam_post/models"
)

type RssParser struct {
	RegParser
}

func (this *RssParser) ParseHtml(post *models.Post) ([]string, error) {
	post.Description = this.RemoveHtml(post.Detail)
	return nil, nil
}

func init() {
	Register("RssParser", &RssParser{})
}
