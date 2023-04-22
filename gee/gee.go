package gee

import (
	"html/template"
	"net/http"
	"strings"
)

type Engine struct {
	*RouterGroup // 可以看做是继承底层模块所拥有的能力
	groups       []*RouterGroup

	htmlTemplates *template.Template // for html render
	funcMap       template.FuncMap   // for html render
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
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

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	// 遍历engine的所有RouterGroup 添加中间层
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}

// gee.go
// Default use Logger() & Recovery middlewares
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
