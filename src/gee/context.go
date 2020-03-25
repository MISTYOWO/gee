package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (context *Context) PostForm(key string) string {
	return context.Req.FormValue(key)
}
func (context *Context) Query(key string) string {
	return context.Req.URL.Query().Get(key)
}
func (context *Context) SetStatus(code int) {
	context.StatusCode = code
	context.Writer.WriteHeader(code)
}
func (context *Context) SetHeader(key string, val string) {
	context.Writer.Header().Set(key, val)
}
func (context *Context) String(code int, format string, vals ...interface{}) {
	context.SetHeader("Content-Type", "text/plain")
	context.SetStatus(code)
	context.Writer.Write([]byte(fmt.Sprintf(format, vals...)))
}
func (context *Context) Json(code int, obj interface{}) {
	context.SetStatus(code)
	context.SetHeader("Content-Type", "application/json")
	encode := json.NewEncoder(context.Writer)
	if err := encode.Encode(obj); err != nil {
		panic(err)
	}

}
func (context *Context) Date(code int, data []byte) {
	context.SetStatus(code)
	context.Writer.Write(data)
}
func (context *Context) Html(code int, html string) {
	context.SetHeader("Content-Type", "text/html")
	context.SetStatus(code)
	context.Writer.Write([]byte(html))
}
