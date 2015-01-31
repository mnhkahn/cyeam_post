package main

import (
	"cyeam_post/app"
)

func main() {
	app.Run()
}

func init() {
	// app.Add("http://localhost:4000/", "CyBot")
	app.Add("http://www.google.com/doodles/doodles.xml", "DoodleBot")
}

// func timer() {
// 	duration := conf.DefaultInt("parse.duration", 60)
// 	timer := time.NewTicker(time.Duration(duration) * time.Minute)
// 	for {
// 		select {
// 		case <-timer.C:
// 			go func() {
// 				Process()
// 			}()
// 		}
// 	}
// }
