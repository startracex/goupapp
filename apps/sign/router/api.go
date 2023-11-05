package router

import (
	"github.com/startracex/goup"
	"goupapp/apps/sign/handler/api"
)

func ApiRoute(group *goup.RouterGroup) {
	group.GET("/user", api.User)
	group.GET("/auth", api.Auth)
}
