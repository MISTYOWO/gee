package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*trieNode
}

func NewRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*trieNode),
	}
}

func ParsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router %s -%s", method, pattern)
	parts := ParsePattern(pattern)
	key := method + ":" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &trieNode{}
	}
	r.roots[method].insertNode(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRouter(method string, path string) (*trieNode, map[string]string) {
	parts := ParsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}
	n := root.searchNode(parts, 0)
	if n != nil {
		partss := ParsePattern(n.pattern)
		for index, part := range partss {
			if part[0] == ':' {
				params[part[1:]] = parts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(parts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
func (r *router) Handle(context *Context) {
	n, params := r.getRouter(context.Method, context.Path)
	if n != nil {
		context.Params = params
		key := context.Method + ":" + n.pattern
		context.handlers = append(context.handlers, r.handlers[key])
	} else {
		context.handlers = append(context.handlers, func(context *Context) {
			context.String(http.StatusNotFound, "404 not found %s\n", context.Path)
		})
	}
	context.Next()
}
