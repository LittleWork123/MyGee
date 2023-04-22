package gee

import (
	"log"
)

type RouterGroup struct {
	*router
	prefix      string
	parent      *RouterGroup
	middlewares []HandlerFunc
}

// 每一个路由的路径肯定是有根的, 那么可以创建一个默认的根路由分组,新加分组必定是基于根路由分组进行新增的
func newRootGroup() *RouterGroup {
	return &RouterGroup{
		prefix: "",
		router: newRouter(),
	}
}

// Use is defined to add middleware to the group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// 所有分组已经路由相关的管理操作的代码,不应该放置到gee.go
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	prefix = group.prefix + prefix
	newGroup := &RouterGroup{
		prefix: prefix,
		parent: group,
		router: group.router,
	}
	return newGroup
}

// 这里类似面向对象编程一样重写router.addRoute 此时的compare为相对路由
func (group *RouterGroup) addRoute(method string, compare string, handler HandlerFunc) {
	pattern := group.prefix + compare
	log.Printf("Route %4s - %s\n", method, pattern)

	// 调用父类的addRoute
	group.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
