package gee

import "net/http"

type Engine struct {
	*RouterGroup // 可以看做是继承底层模块所拥有的能力
	groups       []*RouterGroup
}

func New() *Engine {
	group := newRootGroup()
	engine := &Engine{
		RouterGroup: group,
		groups:      []*RouterGroup{group},
	}
	return engine
}

// engine想知道自己创建了多少个分组, 这里就应该由它自己来进行统计, 这是engine的职责, 而不是RouterGroup的
func (e *Engine) Group(prefix string) *RouterGroup {
	group := e.RouterGroup.Group(prefix)
	e.groups = append(e.groups, group)
	return group
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.handle(c) // 使用底层模块提供的能力
}
