package main

import (
	"fmt"
	gee "gee/v2"
	"net/http"
)

func main() {
	g := gee.New()
	g.GET("/ping", func(ctx *gee.Context) {
		fmt.Println("hey")
		ctx.String(http.StatusOK, "%s", "pong")
	})
	g.Run(":8080")
}