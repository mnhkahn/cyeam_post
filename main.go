package main

import (
	"cyeam_post/dao"
	"cyeam_post/parser"
)

func main() {
	Dao, _ := dao.NewDao("db")
	P, _ := parser.NewParser("cyeam_blog")
	for i := 0; i < P.Len(); i++ {
		Dao.AddPost(P.Index(i))
	}
}
