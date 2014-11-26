package main

import (
	"cyeam_post/dao"
	"cyeam_post/parser"
	"fmt"
	"time"
)

func Process() {
	fmt.Println("Start parse============", time.Now())
	Dao, _ := dao.NewDao("db")
	P, _ := parser.NewParser("cyeam_blog", "http://blog.cyeam.com/rss.xml")
	for i := 0; i < P.Len(); i++ {
		Dao.AddPost(P.Index(i))
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
