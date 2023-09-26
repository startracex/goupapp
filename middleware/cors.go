package middleware

import (
	"github.com/startracex/goup"
	"github.com/startracex/goup/toolkit"
	"log"
)

var DefaultCors = toolkit.Cors{
	AllowOrigin:  []string{"*"},
	AllowMethod:  []string{"*"},
	AllowHeaders: []string{"*"},
}

// SetCors added in new version
func SetCors(c toolkit.Cors) goup.HandlerFunc {
	return goup.SetCors(c)
}

// Cors added in new version
func Cors(s ...string) goup.HandlerFunc {
	return goup.Cors(s...)
}

func Logger() goup.HandlerFunc {
	return goup.Logger(log.Ltime)
}

func Recovery() goup.HandlerFunc {
	return goup.Recovery()
}
