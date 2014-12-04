package parser

import (
	. "cyeam_post/models"
	"gopkg.in/xmlpath.v1"
	"net/http"
	// "time"
	"fmt"
)

type HtmlParser struct {
}

func (this *HtmlParser) Parse(source string) (ParserContainer, error) {
	c := new(HtmlParserContainer)
	c.source = source

	client := http.Client{}
	req, err := http.NewRequest("GET", "http://cyeam.com", nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	path := xmlpath.MustCompile("/html/head/title")
	root, err := xmlpath.Parse(resp.Body)
	if err != nil {
		return c, err
	}
	if value, ok := path.String(root); ok {
		fmt.Println("Found:", value)
	}

	return c, nil
}

type HtmlParserContainer struct {
	source string
	data   []Post
}

func (this *HtmlParserContainer) Index(i int) *Post {
	if len(this.data) > i {
		return &this.data[i]
	}
	return nil
}

func (this *HtmlParserContainer) Set(i int, p *Post) *Post {
	if len(this.data) > i {
		this.data[i] = *p
		return p
	}
	return nil
}

func (this *HtmlParserContainer) Len() int {
	return len(this.data)
}

func init() {
	Register("html", &HtmlParser{})
}
