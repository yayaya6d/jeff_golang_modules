package gin_helper

import "github.com/gin-gonic/gin"

type Router interface {
	SetRouter(routerGroup *gin.RouterGroup)
}

type router struct {
	path        string
	middlewares []gin.HandlerFunc
	subRouters  []Router
	handler     Handler
}

func NewRouter(path string, middlewares []gin.HandlerFunc, subRouter []Router, handler Handler) Router {
	return &router{
		path,
		middlewares,
		subRouter,
		handler,
	}
}

func (r *router) setHandler(routerGroup *gin.RouterGroup) {
	if r.handler.DELETE != nil {
		routerGroup.DELETE(r.path, r.handler.DELETE)
	}
	if r.handler.POST != nil {
		routerGroup.POST(r.path, r.handler.POST)
	}
	if r.handler.GET != nil {
		routerGroup.GET(r.path, r.handler.GET)
	}
	if r.handler.PUT != nil {
		routerGroup.PUT(r.path, r.handler.PUT)
	}
	if r.handler.PATCH != nil {
		routerGroup.PATCH(r.path, r.handler.PATCH)
	}
	if r.handler.OPTIONS != nil {
		routerGroup.OPTIONS(r.path, r.handler.OPTIONS)
	}
}

func (r *router) SetRouter(routerGroup *gin.RouterGroup) {
	r.setHandler(routerGroup)
	for _, middleware := range r.middlewares {
		routerGroup.Use(middleware)
	}

	for _, subRouter := range r.subRouters {
		subRouter.SetRouter(routerGroup.Group(r.path))
	}
}

func RunRouter(router Router, ginEngine *gin.Engine, host string) error {
	router.SetRouter(ginEngine.Group("/"))
	return ginEngine.Run(host)
}
