package main

import (
	"fmt"
	"github.com/startracex/goup"
	"goupapp/apps/web/conf"
	"goupapp/internal/middleware"
)

func main() {
	fmt.Printf("goup@%s: ", goup.Version)
	engine := goup.New()
	engine.Use(middleware.Logger(), middleware.Recovery())
	engine.GET("/", func(request goup.Request, response goup.Response) {
		name := request.GetQuery("name")
		if name == "" {
			name = "World"
		}
		fmt.Fprintf(response.Writer, "<h1>Hello, %s</h1>\n", name)
	})
	engine.GET("/api", func(request goup.Request, response goup.Response) {
		name := request.GetQuery("name")
		if name == "" {
			name = "World"
		}
		response.JSON(map[string]string{
			"name": name,
		})
	})
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
