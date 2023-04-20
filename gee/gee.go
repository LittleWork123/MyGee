package gee

import (
	"fmt"
	"log"
	"net/http"
)

// 定义处理Request的句柄函数
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	routerGroup map[string]HandlerFunc
}

// New 一个Engine
func New() *Engine {
	return &Engine{routerGroup: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, Handler HandlerFunc) {
	key := method + "-" + pattern
	if _, existed := engine.routerGroup[key]; existed {
		log.Fatal("the method and pattern you added already existed")
	}
	engine.routerGroup[key] = Handler
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if Handler, ok := engine.routerGroup[key]; ok {
		Handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND. %s \n", req.URL)
	}
}

// GET 方法调用
func (engine *Engine) GET(pattern string, Handler HandlerFunc) {
	engine.addRoute("GET", pattern, Handler)
}

// POST 方法调用
func (engine *Engine) POST(pattern string, Handler HandlerFunc) {
	engine.addRoute("POST", pattern, Handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
