package router

import (
	"gen-concept-api/api/handler"
	"gen-concept-api/config"

	"github.com/gin-gonic/gin"
)

func Importer(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewImporterHandler(cfg)

	r.POST("/parse", h.Parse)
}
