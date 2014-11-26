package main

import (
	"cyeam_post/dao"
	"cyeam_post/parser"
	"fmt"
	"time"
)

func Process() {
	fmt.Println("Start parse============", time.Now())
	Dao, _ := dao.NewDao("db", "cyeam:qwerty@tcp(128.199.131.129:3306)/cyeam?charset=utf8 ")
	P, _ := parser.NewParser("cyeam_blog", "http://blog.cyeam.com/rss.xml")
	for i := 0; i < P.Len(); i++ {
		if Dao.IsPostUpdate(P.Index(i)) {
			Dao.UpdatePost(P.Index(i))
		} else if Dao.GetPostByLink(P.Index(i).Link) == nil {
			Dao.AddPost(P.Index(i))
		}
	}
}

func timer() {
	timer := time.NewTicker(2 * time.Minute)
	for {
		select {
		case <-timer.C:
			go func() {
				Process()
			}()
		}
	}
}

func main() {
	// timer()
	Process()
}
