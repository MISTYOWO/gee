package gee

import (
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := NewRouter()
	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/hello/:name", nil)
	r.addRouter("GET", "/hello/b/c", nil)
	r.addRouter("GET", "/hi/:name", nil)
	r.addRouter("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(ParsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(ParsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(ParsePattern("/p/*name/*"), []string{"p", "*name"})

	if !ok {
		t.Fatal("test parsePattern is fail")
	}
}
func TestGetRouter(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRouter("GET", "/hello/geekutu")
	if n == nil {
		t.Fatal("GET /hello/geekutu fails")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("/hello/:name not match")
	}
	if ps["name"] != "geekutu" {
		t.Fatal("should be geekutu")
	}
}
