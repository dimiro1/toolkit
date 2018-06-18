package router

import (
	"net/http"
)

type Router interface {
	http.Handler

	Get(path string, handler http.HandlerFunc)
	Post(path string, handler http.HandlerFunc)
	Put(path string, handler http.HandlerFunc)
	Delete(path string, handler http.HandlerFunc)
	Head(path string, handler http.HandlerFunc)
	Connect(path string, handler http.HandlerFunc)
	Options(path string, handler http.HandlerFunc)
	Patch(path string, handler http.HandlerFunc)
	Trace(path string, handler http.HandlerFunc)

	Handle(method string, path string, handler http.HandlerFunc)
	NotFound(handler http.HandlerFunc)
}
