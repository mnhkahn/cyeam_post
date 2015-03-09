package dao

import (
	. "cyeam_post/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/franela/goreq"
	"net/url"
)

type DuoShuoConfig struct {
	ShortName string `json:"short_name"`
	Secret    string `json:"secret"`
}

type DuoShuoDao struct {
}

func (this *DuoShuoDao) NewDaoImpl(dsn string) (DaoContainer, error) {
	d := new(DuoShuoDaoContainer)

	config := new(DuoShuoConfig)
	err := json.Unmarshal([]byte(dsn), config)
	d.config = config
	if err != nil {
		return d, fmt.Errorf("Config for duoshuo is error: %v", err)
	}

	return d, nil
}

type DuoShuoDaoContainer struct {
	config   *DuoShuoConfig
	is_debug bool
	req      goreq.Request
}

func (this *DuoShuoDaoContainer) Debug(is_debug bool) {
	this.req.ShowDebug = is_debug
}

func (this *DuoShuoDaoContainer) AddPost(p *Post) {
	this.req.Method = "POST"
	this.req.Uri = "http://api.duoshuo.com/posts/import.json"
	this.req.ContentType = "application/x-www-form-urlencoded"

	addDuoShuo := url.Values{}
	addDuoShuo.Add("short_name", this.config.ShortName)
	addDuoShuo.Add("secret", this.config.Secret)
	addDuoShuo.Add("posts[0][post_key]", p.Link)
	addDuoShuo.Add("posts[0][thread_key]", "haixiuzucyeam")

	p.Description = ""
	p.Detail = ""
	duoshuo_byte, _ := json.Marshal(*p)
	addDuoShuo.Add("posts[0][message]", base64.URLEncoding.EncodeToString(duoshuo_byte))
	this.req.Body = addDuoShuo.Encode()
	this.req.ShowDebug = true
	resp, err := this.req.Do()
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		err_str, _ := resp.Body.ToString()
		panic(fmt.Errorf("Error: %d, %s", resp.StatusCode, err_str))
	}
}

func (this *DuoShuoDaoContainer) AddPosts(p []Post) {

}

func (this *DuoShuoDaoContainer) DelPost(id interface{}) {

}

func (this *DuoShuoDaoContainer) DelPosts(source string) {

}

func (this *DuoShuoDaoContainer) UpdatePost(p *Post) {

}

func (this *DuoShuoDaoContainer) AddOrUpdate(p *Post) {
	this.AddPost(p)
}

func (this *DuoShuoDaoContainer) GetPostById(id int) *Post {
	p := new(Post)
	return p
}

func (this *DuoShuoDaoContainer) GetPostByLink(url string) *Post {
	p := new(Post)
	return p
}

func (this *DuoShuoDaoContainer) GetPost(author, sort string, limit, start int) []Post {
	return nil
}

func (this *DuoShuoDaoContainer) IsPostUpdate(p *Post) bool {
	is_update := false
	return is_update
}

func (this *DuoShuoDaoContainer) Search(q string, limit, start int) (int, float64, []Post) {
	println("*************8serach")
	this.req.Method = "GET"
	this.req.Uri = "http://api.duoshuo.com/threads/listPosts.json"
	this.req.ContentType = "application/x-www-form-urlencoded"

	// addDuoShuo := url.Values{}
	// addDuoShuo.Add("short_name", this.config.ShortName)
	// addDuoShuo.Add("secret", this.config.Secret)
	// addDuoShuo.Add("posts[0][post_key]", p.Id)
	// addDuoShuo.Add("posts[0][thread_key]", "haixiuzu-cyeam")

	// duoshuo_byte, _ := json.Marshal(addDuoShuo)
	// addDuoShuo.Add("posts[0][message]", base64.URLEncoding.EncodeToString(duoshuo_byte))
	// this.req.Body = addDuoShuo.Encode()

	// resp, err := this.req.Do()
	// if err != nil {
	// 	panic(err)
	// }
	// if resp.StatusCode != 200 {
	// 	err_str, _ := resp.Body.ToString()
	// 	panic(err_str)
	// }
	return 0, 0, nil
}

func init() {
	Register("duoshuo", &DuoShuoDao{})
}
