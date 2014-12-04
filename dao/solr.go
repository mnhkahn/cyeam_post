package dao

import (
	. "cyeam_post/models"
	"github.com/franela/goreq"
)

type SolrDao struct {
	Dsn      string
	is_debug bool
}

func (this *SolrDao) NewDaoImpl(dsn string) (SolrDaoContainer, error) {
	solr := new(SolrDaoContainer)
	solr.solr_req = goreq.Request{
		Timeout:     time.Duration(5) * time.Second,
		Compression: goreq.Gzip(),
	}
	solr.solr_req.SetConnectTimeout(time.Duration(5) * time.Second)
	return solr, nil
}

type SolrDaoContainer struct {
	solr_req goreq.Request
}

func (this *SolrDaoContainer) Debug(is_debug bool) {
	this.is_debug = is_debug
}

func (this *SolrDaoContainer) AddPost(p *Post) {

}

func (this *SolrDaoContainer) AddPosts(p []Post) {

}

func (this *SolrDaoContainer) DelPost(id int) {

}

func (this *SolrDaoContainer) DelPosts(source string) {

}

func (this *SolrDaoContainer) UpdatePost(p *Post) {

}

func (this *SolrDaoContainer) GetPostById(id int) *Post {
	p := new(Post)
	return p
}

func (this *SolrDaoContainer) GetPost(author, sort string, limit, start int) []Post {
	p := make([]Post, 0)
	return p
}

func (this *SolrDaoContainer) IsPostUpdate(p *Post) bool {
	is_update := false
	return is_update
}

func (this *DbDaoContainer) Search(q string, limit, start int) []Post {
	res := make([]Post, 0)
	return res
}

func init() {
	Register("solr", &SolrDao{})
}
