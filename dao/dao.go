package dao

import (
	. "cyeam_post/models"
	"fmt"
)

type Dao interface {
	AddPost(p *Post)
	AddPosts(p []Post)
	DelPost(id int)
	DelPosts(source string)
	UpdatePost(p *Post)
	GetPostById(id int) *Post
	Search(q string) []Post
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

func NewDao(dao_name string) (Dao, error) {
	dao, ok := daos[dao_name]
	if !ok {
		return nil, fmt.Errorf("parser: unknown dao_name %q", dao_name)
	}
	return dao, nil
}
