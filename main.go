// main project main.go
package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.Get("/index", func(c *gee.Context) {
		c.Html(http.StatusOK, "<h1>index page</h1>")
	})

	v1 := r.NewGroup("/v1")
	v1.Use(gee.LoggerB())
	v1.Get("/", func(c *gee.Context) {
		c.Html(http.StatusOK, "<h1>hello Gee</h1>")
	})
	v1.Get("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	v2 := r.NewGroup("/v2")

	v2.Get("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,you are at %s\n", c.Param("name"), c.Path)
	})
	v2.Post("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
