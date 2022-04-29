package app

import (
	"go-server/internal/app/middleware"
	"go-server/internal/app/router"

	"github.com/gin-gonic/gin"
)

func InitGin(r *router.Router) *gin.Engine {
	cfg := GetConfig()
	gin.SetMode(cfg.RunMode)
	app := gin.New()
	app.Use(middleware.RecoveryMiddleware())
	r.RegisterAPI(app)

	return app
}
