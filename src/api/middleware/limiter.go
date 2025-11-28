package middleware

import (
	"net/http"

	"gen-concept-api/api/helper"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func LimitByRequest() gin.HandlerFunc {
	// Increased from 1 to 100 requests per second for development
	lmt := tollbooth.NewLimiter(100, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(nil, false, helper.LimiterError, err))
			return
		} else {
			c.Next()
		}
	}
}
