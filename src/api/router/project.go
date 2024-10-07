package router

import (
	"gen-concept-api/api/handler"
	"gen-concept-api/config"

	"github.com/gin-gonic/gin"
)

func Project(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewProjectHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}

