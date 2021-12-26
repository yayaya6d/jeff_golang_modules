package gin_helper

import "github.com/gin-gonic/gin"

type Handler struct {
	GET     gin.HandlerFunc
	POST    gin.HandlerFunc
	DELETE  gin.HandlerFunc
	PATCH   gin.HandlerFunc
	PUT     gin.HandlerFunc
	OPTIONS gin.HandlerFunc
}
