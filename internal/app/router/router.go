package router

import (
	"go-server/internal/app/admin/api"

	"github.com/gin-gonic/gin"
)

type Router struct {
	UserApi *api.UserAPi
}

func (r *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	// g.Use()

	v1 := g.Group("/v1")
	{
		{
			v1.GET("/users", r.UserApi.List)
			v1.GET("/users/:userId", r.UserApi.Retrieve)
			v1.POST("/users", r.UserApi.Create)
			v1.POST("/bulk-users", r.UserApi.BulkCreate)
			v1.PUT("/users/:userId", r.UserApi.Update)
			v1.DELETE("/users/:userId", r.UserApi.Delete)
		}
	}
}
