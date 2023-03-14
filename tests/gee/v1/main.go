package main

import (
	gee "gee/v1"
	"log"
	"net/http"
)

func main() {
	g := gee.New()
	g.GET("/ping", func (w http.ResponseWriter, req *http.Request){
		log.Print("someone pinged")
		w.Write([]byte("pong"))
	})
}