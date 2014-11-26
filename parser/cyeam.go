package parser

import (
	. "cyeam_post/models"
)

type CyeamBlogParser struct {
}

func (this *CyeamBlogParser) Parse() (ParserContainer, error) {
	c := new(CyeamBlogParserContainer)
	return c, nil
}

type CyeamBlogParserContainer struct {
	source string
	data   []Post
}

func (this *CyeamBlogParserContainer) Index(i int) *Post {
	if len(this.data) > i {
		return &this.data[i]
	}
	return nil
}

func (this *CyeamBlogParserContainer) Set(i int, p *Post) *Post {
	if len(this.data) > i {
		this.data[i] = *p
		return p
	}
	return nil
}

func (this *CyeamBlogParserContainer) Len() int {
	return len(this.data)
}
func init() {
	Register("cyeam_blog", &CyeamBlogParser{})
}
