package dao

// import (
// 	. "cyeam_post/models"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/go-xorm/xorm"
// )

// type DbDao struct {
// 	Dsn string
// }

// func (this *DbDao) NewDaoImpl(dsn string) (DaoContainer, error) {
// 	db := new(DbDaoContainer)
// 	var err error
// 	db.Engine, err = xorm.NewEngine("mysql", dsn)
// 	if err != nil {
// 		return db, err
// 	}
// 	db.Engine.SetMaxConns(5)
// 	return db, nil
// }

// type DbDaoContainer struct {
// 	Engine *xorm.Engine
// }

// func (this *DbDaoContainer) Debug(is_debug bool) {
// 	this.Engine.ShowDebug = is_debug
// }

// func (this *DbDaoContainer) AddPost(p *Post) {
// 	this.Engine.Table("post").Insert(p)
// }

// func (this *DbDaoContainer) AddPosts(p []Post) {

// }

// func (this *DbDaoContainer) DelPost(id int) {

// }

// func (this *DbDaoContainer) DelPosts(source string) {

// }

// func (this *DbDaoContainer) UpdatePost(p *Post) {
// 	this.Engine.Table("post").Where("link=?", p.Link).Update(p)
// }

// func (this *DbDaoContainer) GetPostById(id int) *Post {
// 	p := new(Post)
// 	p.Id = id
// 	this.Engine.Table("post").Get(p)
// 	return p
// }

// func (this *DbDaoContainer) GetPostByLink(url string) *Post {
// 	p := new(Post)
// 	p.Link = url
// 	_, err := this.Engine.Table("post").Get(p)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return p
// }

// func (this *DbDaoContainer) GetPost(author, sort string, limit, start int) []Post {
// 	p := make([]Post, 0)
// 	if author == "" {
// 		author = "TRUE"
// 	}
// 	if sort == "" {
// 		sort = "create_time desc"
// 	}
// 	this.Engine.Table("post").Where("author=?", author).OrderBy(sort).Limit(limit, start).Find(&p)
// 	return p
// }

// func (this *DbDaoContainer) IsPostUpdate(p *Post) bool {
// 	is_update := false
// 	temp := this.GetPostByLink(p.Link)
// 	if temp.Title != "" {
// 		if temp.Title != p.Title || temp.Author != p.Author || temp.Detail != p.Detail || temp.Figure != p.Figure {
// 			is_update = true
// 		}
// 	}
// 	return is_update
// }

// func (this *DbDaoContainer) Search(q string, limit, start int) []Post {
// 	res := make([]Post, 0)
// 	if q == "" {
// 		q = "%"
// 	} else {
// 		q = "%" + q + "%"
// 	}
// 	this.Engine.ShowDebug = true
// 	err := this.Engine.Table("post").Decr("create_time").Limit(limit, start).Where("title like ? OR detail like ?", q, q).Find(&res)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return res
// }

// func init() {
// 	Register("db", &DbDao{})
// }
