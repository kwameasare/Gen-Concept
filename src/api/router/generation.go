package router

import (
	"gen-concept-api/api/handler"
	"gen-concept-api/config"

	"github.com/gin-gonic/gin"
)

func Generation(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewGenerationHandler(cfg)

	r.POST("/preview", h.Preview)
}
