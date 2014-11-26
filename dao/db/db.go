package db

import (
	. "cyeam_post"
	. "cyeam_post/models"
)

type DbDaoImpl struct {
}

func (this *DbDaoImpl) AddPost(p *Post) {

}

func (this *DbDaoImpl) AddPosts(p []Post) {

}

func (this *DbDaoImpl) DelPost(id int) {

}

func (this *DbDaoImpl) DelPosts(source string) {

}

func (this *DbDaoImpl) UpdatePost(p *Post) {

}

func (this *DbDaoImpl) GetPostById(id int) *Post {
	return nil
}

func (this *DbDaoImpl) Search(q string) []Post {
	res := make([]Post, 0)
	return res
}

func init() {
	Register("db", &DbDaoImpl{})
}
