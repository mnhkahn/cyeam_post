package parser

import (
	"cyeam_post/models"
)

type RssParser struct {
	RegParser
	is_debug bool
}

func (this *RssParser) Debug(is_debug bool) {
	this.is_debug = is_debug
}

func (this *RssParser) ParseHtml(post *models.Post) ([]string, error) {
	post.Description = this.RemoveHtml(post.Detail)
	return nil, nil
}

func init() {
	Register("RssParser", &RssParser{})
}
