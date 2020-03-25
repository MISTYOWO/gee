// main project main.go
package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.Html(http.StatusOK, "<h1>hello gee</h1>")
	})
	r.Get("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,you are at %S\n", c.Query("name"), c.Path)
	})
	r.Post("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
