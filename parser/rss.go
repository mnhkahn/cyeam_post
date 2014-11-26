package parser

import (
	. "cyeam_post/models"
	"encoding/xml"
	"github.com/astaxie/beego/httplib"
	"strings"
	"time"
)

type RssParser struct {
}

func (this *RssParser) Parse(source string) (ParserContainer, error) {
	c := new(RssParserContainer)
	c.source = source

	res := RssFeed{}
	req := httplib.Get(c.source)
	err := req.ToXml(&res)
	if err != nil {
		return c, err
	}
	for _, item := range res.Channel.Items {
		post := Post{}
		post.Title = item.Title
		post.CreateTime, err = time.Parse(time.RFC3339, strings.Replace(item.PubDate, "+", "Z", 1))
		if err != nil {
			post.CreateTime = time.Now()
		}
		post.Author = item.Author
		post.Detail = item.Description
		post.Category = item.Category
		post.Figure = item.Figure
		c.data = append(c.data, post)
	}
	return c, nil
}

type RssParserContainer struct {
	source string
	data   []Post
}

func (this *RssParserContainer) Index(i int) *Post {
	if len(this.data) > i {
		return &this.data[i]
	}
	return nil
}

func (this *RssParserContainer) Set(i int, p *Post) *Post {
	if len(this.data) > i {
		this.data[i] = *p
		return p
	}
	return nil
}

func (this *RssParserContainer) Len() int {
	return len(this.data)
}
func init() {
	Register("cyeam_blog", &RssParser{})
}

type RssFeed struct {
	XMLName xml.Name    `xml:"rss"`
	Channel *RssChannel `xml:"channel"`
}

type RssChannel struct {
	XMLName       xml.Name   `xml:"channel"`
	Title         string     `xml:"title"`
	Description   string     `xml:"description"`
	Link          string     `xml:"link"`
	Language      string     `xml:"language"`
	PubDate       string     `xml:"pubDate"`
	LastBuildDate string     `xml:"lastBuildDate"`
	Items         []*RssItem `xml:"item"`
}

type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"` // required
	Figure      string   `xml:"figure"`
	Link        string   `xml:"link"`        // required
	Description string   `xml:"description"` // required
	Author      string   `xml:"author,omitempty"`
	Category    string   `xml:"category,omitempty"`
	Comments    string   `xml:"comments,omitempty"`
	Guid        string   `xml:"guid,omitempty"`    // Id used
	PubDate     string   `xml:"pubDate,omitempty"` // created or updated
	Source      string   `xml:"source,omitempty"`
}
