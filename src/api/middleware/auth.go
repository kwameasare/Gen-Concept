package middleware

import (
	"net/http"
	"strings"

	"gen-concept-api/api/helper"
	"gen-concept-api/config"
	constant "gen-concept-api/constant"
	"gen-concept-api/pkg/logging"
	"gen-concept-api/pkg/service_errors"
	"gen-concept-api/usecase"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenUsecase = usecase.NewTokenUsecase(cfg)

	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := c.GetHeader(constant.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if auth == "" || len(token) < 2 {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		} else {
			claimMap, err = tokenUsecase.GetClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				}
			}
		}
		if err != nil {
			// Ensure CORS headers are set even when authentication fails
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(
				nil, false, helper.AuthError, err,
			))
			return
		}

		c.Set(constant.UserIdKey, claimMap[constant.UserIdKey])
		c.Set(constant.FirstNameKey, claimMap[constant.FirstNameKey])
		c.Set(constant.LastNameKey, claimMap[constant.LastNameKey])
		c.Set(constant.UsernameKey, claimMap[constant.UsernameKey])
		c.Set(constant.EmailKey, claimMap[constant.EmailKey])
		c.Set(constant.MobileNumberKey, claimMap[constant.MobileNumberKey])
		c.Set(constant.RolesKey, claimMap[constant.RolesKey])
		c.Set(constant.ExpireTimeKey, claimMap[constant.ExpireTimeKey])

		c.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	logger := logging.NewLogger(config.GetConfig())
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			logger.Warn(logging.Validation, logging.Api, "Authorization failed: no context keys", nil)
			// Ensure CORS headers are set even when authorization fails
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, helper.ForbiddenError))
			return
		}
		rolesVal := c.Keys[constant.RolesKey]
		// fmt.Println(rolesVal)
		if rolesVal == nil {
			logger.Warn(logging.Validation, logging.Api, "Authorization failed: no roles in context", nil)
			// Ensure CORS headers are set even when authorization fails
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, helper.ForbiddenError))
			return
		}
		roles := rolesVal.([]interface{})
		val := map[string]int{}
		for _, item := range roles {
			val[item.(string)] = 0
		}

		for _, item := range validRoles {
			if _, ok := val[item]; ok {
				c.Next()
				return
			}
		}

		logger.Warn(logging.Validation, logging.Api, "Authorization failed: role mismatch", map[logging.ExtraKey]interface{}{
			"UserRoles":      roles,
			"RequiredRoles":  validRoles,
			logging.Username: c.Keys[constant.UsernameKey],
		})

		// Ensure CORS headers are set even when authorization fails
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

		c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, helper.ForbiddenError))
	}
}
