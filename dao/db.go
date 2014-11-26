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

}

func (this *DbDaoContainer) GetPostById(id int) *Post {
	return nil
}

func (this *DbDaoContainer) Search(q string) []Post {
	res := make([]Post, 0)
	return res
}

func init() {
	Register("db", &DbDao{})
}
