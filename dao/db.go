package dao

import (
	. "cyeam_post/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type DbDao struct {
	Dsn string
}

func (this *DbDao) NewDaoImpl(dsn string) (DaoContainer, error) {
	db := new(DbDaoContainer)
	var err error
	db.Engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return db, err
	}
	db.Engine.SetMaxConns(5)
	return db, nil
}

type DbDaoContainer struct {
	Engine *xorm.Engine
}

func (this *DbDaoContainer) AddPost(p *Post) {
	this.Engine.Table("post").Insert(p)
}

func (this *DbDaoContainer) AddPosts(p []Post) {

}

func (this *DbDaoContainer) DelPost(id int) {

}

func (this *DbDaoContainer) DelPosts(source string) {

}

func (this *DbDaoContainer) UpdatePost(p *Post) {
	this.Engine.Table("post").Update(p)
}

func (this *DbDaoContainer) GetPostById(id int) *Post {
	p := new(Post)
	p.Id = id
	this.Engine.Table("post").Get(p)
	return p
}

func (this *DbDaoContainer) GetPostByLink(url string) *Post {
	p := new(Post)
	p.Link = url
	this.Engine.Table("post").Get(p)
	return p
}
func (this *DbDaoContainer) IsPostUpdate(p *Post) bool {
	is_exists := false
	temp := this.GetPostByLink(p.Link)
	if temp != nil {
		if temp.Title != p.Title || temp.Author != p.Author || temp.Detail != p.Detail {
			is_exists = true
		}
	}
	return is_exists
}

func (this *DbDaoContainer) Search(q string) []Post {
	res := make([]Post, 0)
	return res
}

func init() {
	Register("db", &DbDao{})
}
