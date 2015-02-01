package parser

import (
	"cyeam_post/models"
	"encoding/xml"
	"github.com/astaxie/beego/httplib"
	"strings"
	"time"
)

type RssParser struct {
	RegParser
	is_debug bool
}

func (this *RssParser) Debug(is_debug bool) {
	this.is_debug = is_debug
}

func (this *RssParser) ParseHtml(post *models.Post) ([]string, error) {
	res := RssFeed{}
	req := httplib.Get(post.Link)
	req.SetTimeout(5*time.Second, 5*time.Second)
	err := req.ToXml(&res)
	if err != nil {
		panic(err)
	}
	for i, item := range res.Channel.Items {
		if i > 1 {
			break
		}
		post.Title = item.Title
		if item.PubDate == "" {
			post.CreateTime.Time = time.Now()
		} else if temp_date, err := time.Parse("2006-01-02T15:04:05", string([]byte(item.PubDate)[:19])); err == nil {
			post.CreateTime.Time = temp_date
		} else if temp_date, err := time.Parse("2006-01-02 15:04:05", item.PubDate); err == nil {
			post.CreateTime.Time = temp_date
		} else {
			post.CreateTime.Time = time.Now()
		}
		if item.Author != "" {
			post.Author = item.Author
		} else {
			post.Author = res.Channel.Title
		}
		post.Detail = item.Description
		post.Category = item.Category

		temp_document, err := NewCssParser(strings.NewReader(item.Description))
		if err != nil {
			panic(err)
		}
		imgs := temp_document.GetImgs()
		if len(imgs) > 0 {
			imgs[0] = strings.TrimLeft(imgs[0], "/")
			if strings.HasPrefix(imgs[0], "http") {
				post.Figure = imgs[0]
			} else {
				post.Figure = "http://" + imgs[0]
			}
		}

		post.Source = item.Link
		post.Link = item.Link
		post.ParseDate.Time = time.Now()
		post.Description = this.RemoveHtml(post.Detail)
	}
	return nil, nil
}

func init() {
	Register("RssParser", &RssParser{})
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
