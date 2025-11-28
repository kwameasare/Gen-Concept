package router

import (
	"gen-concept-api/api/handler"
	"gen-concept-api/config"

	"github.com/gin-gonic/gin"
)

func Blueprint(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewBlueprintHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}
