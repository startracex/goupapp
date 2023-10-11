package main

import (
	"fmt"
	"github.com/startracex/goup"
	"goupapp/apps/sign/conf"
	"goupapp/apps/sign/router"
	"goupapp/internal/middleware"
)

func main() {
	fmt.Printf("goup@%s: ", goup.Version)
	engine := goup.New()
	engine.Use(middleware.Recovery(), middleware.SetCors(middleware.DefaultCors))
	engine.LoadHTMLFiles("./client/index.html")
	for _, pattern := range []string{"/", "login", "/signup", "/aith"} {
		engine.GET(pattern, func(request goup.Request, response goup.Response) {
			response.HTML("index.html", nil)
		})
	}
	engine.Public("/public", "./client/public")
	router.Init(engine)
	var err error
	if conf.TLS {
		fmt.Printf("Listen and serve on https://localhost:%s\n", conf.PORT)
		err = engine.ListenAndServe(conf.PORT)
	} else {
		fmt.Printf("Listen and serve on http://localhost:%s\n", conf.PORT)
		err = engine.ListenAndServe(conf.PORT)
	}
	if err != nil {
		fmt.Println(err)
	}
}
