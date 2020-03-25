package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}
func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router %s -%s", method, pattern)
	key := method + ":" + pattern
	r.handlers[key] = handler
}
func (r *router) Handle(context *Context) {
	key := context.Method + ":" + context.Path
	if handler, ok := r.handlers[key]; ok {
		handler(context)
	} else {
		context.String(http.StatusNotFound, "404 not found %s\n", context.Path)
	}
}
