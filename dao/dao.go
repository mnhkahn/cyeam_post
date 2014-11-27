package dao

import (
	. "cyeam_post/models"
	"fmt"
)

type DaoContainer interface {
	AddPost(p *Post)
	AddPosts(p []Post)
	DelPost(id int)
	DelPosts(source string)
	UpdatePost(p *Post)
	GetPostById(id int) *Post
	GetPostByLink(url string) *Post
	GetPost(author, sort string, limit, start int) []Post
	IsPostUpdate(p *Post) bool
	Search(q string) []Post
	Debug(is_debug bool)
}

type Dao interface {
	NewDaoImpl(dsn string) (DaoContainer, error)
}

var daos = make(map[string]Dao)

// Register makes a config adapter available by the adapter name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, dao Dao) {
	if dao == nil {
		panic("dao: Register dao is nil")
	}
	if _, ok := daos[name]; ok {
		panic("dao: Register called twice for adapter " + name)
	}
	daos[name] = dao
}

func NewDao(dao_name, dsn string) (DaoContainer, error) {
	dao, ok := daos[dao_name]
	if !ok {
		return nil, fmt.Errorf("parser: unknown dao_name %q", dao_name)
	}
	return dao.NewDaoImpl(dsn)
}
