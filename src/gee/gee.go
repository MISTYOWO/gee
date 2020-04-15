// Gee project Gee.go
package gee

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
	groups []*RouterGroup
	*RouterGroup
	httpTemplates *template.Template
	funcMap       template.FuncMap
}

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	engine.router.addRouter(method, pattern, handler)
}

func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)

	var midwares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			midwares = append(midwares, group.middlewares...)
		}
	}
	c.handlers = append(c.handlers, midwares...)
	c.engine = engine
	engine.router.Handle(c)
}

func (engine *Engine) SetFunMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}
func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.httpTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

func Logger() HandlerFunc {
	return func(context *Context) {
		log.Printf("Logger start")
		context.Next()
		log.Printf("Logger end")
	}
}

func LoggerB() HandlerFunc {
	return func(context *Context) {
		log.Printf("LoggerB start")
		context.Next()
		log.Printf("LoggerB end")
	}
}
