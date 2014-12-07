package dao

import (
	. "cyeam_post/models"
	// "encoding/json"
	"fmt"
	"github.com/franela/goreq"
	"net/url"
	"time"
)

type SolrDao struct {
	Dsn string
}

func (this *SolrDao) NewDaoImpl(dsn string) (DaoContainer, error) {
	solr := new(SolrDaoContainer)
	solr.dsn = dsn
	solr.solr_req = goreq.Request{
		ContentType: "application/json",
		UserAgent:   "Cyeam",
		Timeout:     time.Duration(5) * time.Second,
		// Compression: goreq.Gzip(),
	}
	solr.solr_req.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
	goreq.SetConnectTimeout(time.Duration(5) * time.Second)
	return solr, nil
}

type SolrDaoContainer struct {
	dsn      string
	is_debug bool
	solr_req goreq.Request
}

func (this *SolrDaoContainer) Debug(is_debug bool) {
	this.is_debug = is_debug
}

func (this *SolrDaoContainer) AddPost(p *Post) {
	this.solr_req.Method = "POST"
	this.solr_req.Uri = this.dsn + "/update"

	addSolr := new(AddSolr)
	addSolr.Add.CommitWithin = 1000
	addSolr.Add.Doc = *p
	addSolr.Add.Overwrite = true

	query := url.Values{}
	query.Add("wt", "json")
	this.solr_req.Body = *addSolr
	this.solr_req.QueryString = query

	if this.is_debug {
		this.showDebug()
	}
	res, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}

	str, _ := res.Body.ToString()
	fmt.Println(str)
}

func (this *SolrDaoContainer) AddPosts(p []Post) {

}

func (this *SolrDaoContainer) DelPost(id interface{}) {
	this.solr_req.Method = "POST"
	this.solr_req.Uri = this.dsn + "/update"

	delSolr := new(DelSolr)
	delSolr.Del.Query = fmt.Sprintf(`id:%v`, id)
	delSolr.Del.CommitWithin = 1000

	query := url.Values{}
	query.Add("wt", "json")
	this.solr_req.Body = *delSolr
	this.solr_req.QueryString = query

	if this.is_debug {
		this.showDebug()
	}
	res, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}

	str, _ := res.Body.ToString()
	fmt.Println(str)
}

func (this *SolrDaoContainer) DelPosts(source string) {
	this.solr_req.Method = "POST"
	this.solr_req.Uri = this.dsn + "/update"

	delSolr := new(DelSolr)
	delSolr.Del.Query = fmt.Sprintf(`source:%s`, source)
	delSolr.Del.CommitWithin = 1000

	query := url.Values{}
	query.Add("wt", "json")
	this.solr_req.Body = *delSolr
	this.solr_req.QueryString = query

	if this.is_debug {
		this.showDebug()
	}
	res, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}

	str, _ := res.Body.ToString()
	fmt.Println(str)
}

func (this *SolrDaoContainer) UpdatePost(p *Post) {

}

func (this *SolrDaoContainer) GetPostById(id int) *Post {
	p := new(Post)
	return p
}

func (this *SolrDaoContainer) GetPostByLink(url string) *Post {
	p := new(Post)
	return p
}

func (this *SolrDaoContainer) GetPost(author, sort string, limit, start int) []Post {
	this.solr_req.Method = "GET"
	this.solr_req.Uri = this.dsn + "/select"

	query := url.Values{}
	query.Add("wt", "json")
	query.Add("q", fmt.Sprintf("author:%s", author))
	if sort != "" {
		query.Add("sort", sort)
	}
	query.Add("start", fmt.Sprintf("%d", start))
	query.Add("rows", fmt.Sprintf("%d", limit))
	this.solr_req.QueryString = query

	if this.is_debug {
		this.showDebug()
	}
	res, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}

	solr_posts := new(SolrPost)
	err = res.Body.FromJsonTo(solr_posts)
	if err != nil {
		panic(err)
	}
	return solr_posts.Response.Docs
}

func (this *SolrDaoContainer) IsPostUpdate(p *Post) bool {
	is_update := false
	return is_update
}

func (this *SolrDaoContainer) Search(q string, limit, start int) []Post {
	this.solr_req.Method = "GET"
	this.solr_req.Uri = this.dsn + "/select"

	query := url.Values{}
	query.Add("wt", "json")
	query.Add("q", fmt.Sprintf("title:*%s* AND detail:*%s*", q, q))
	query.Add("start", fmt.Sprintf("%d", start))
	query.Add("rows", fmt.Sprintf("%d", limit))
	this.solr_req.QueryString = query

	if this.is_debug {
		this.showDebug()
	}
	res, err := this.solr_req.Do()
	if err != nil {
		panic(err)
	}

	solr_posts := new(SolrPost)
	err = res.Body.FromJsonTo(solr_posts)
	if err != nil {
		panic(err)
	}
	return solr_posts.Response.Docs
}

func (this *SolrDaoContainer) showDebug() {
	if this.is_debug {
		debug_url := this.solr_req.Uri
		if this.solr_req.QueryString != nil {
			debug_url += "?" + url.Values(this.solr_req.QueryString.(url.Values)).Encode()
		}
		fmt.Printf("[solr] %s\n %v\n", debug_url, this.solr_req.Body)
	}
}

func init() {
	Register("solr", &SolrDao{})
}

type AddSolr struct {
	Add struct {
		CommitWithin int  `json:"commitWithin"`
		Doc          Post `json:"doc"`
		Overwrite    bool `json:"overwrite"`
	} `json:"add"`
}

type DelSolr struct {
	Del struct {
		Query        string `json:"query"`
		CommitWithin int    `json:"commitWithin"`
	} `json:"delete"`
}

type SolrPost struct {
	Response struct {
		Docs     []Post  `json:"docs"`
		NumFound float64 `json:"numFound"`
		Start    float64 `json:"start"`
	} `json:"response"`
	ResponseHeader SolrResponseHeader `json:"responseHeader"`
	Error          SolrError          `json:"error"`
}

type SolrResponseHeader struct {
	QTime  float64 `json:"QTime"`
	Params struct {
		Indent string `json:"indent"`
		Q      string `json:"q"`
		Wt     string `json:"wt"`
	} `json:"params"`
	Status float64 `json:"status"`
}

type SolrError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
