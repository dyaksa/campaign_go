package middleware

import (
	"campaignproject/auth"
	"campaignproject/helper"
	"campaignproject/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	IDToken string `header:"Authorization" binding:"required"`
}

type middleware struct {
	authService auth.Service
	userService user.Service
}

func NewMiddleware(authService auth.Service, userService user.Service) *middleware {
	return &middleware{authService: authService, userService: userService}
}

func (m *middleware) AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var h authHeader
		err := c.ShouldBindHeader(&h)
		if err != nil {
			data := gin.H{"is_valid": false}
			responsJSON := helper.APIResponse(err.Error(), http.StatusUnauthorized, "errors", data)
			c.AbortWithStatusJSON(http.StatusUnauthorized, responsJSON)
			return
		}
		ok := strings.Contains(h.IDToken, "Bearer")
		if !ok {
			data := gin.H{"is_valid": false}
			responseJSON := helper.APIResponse("jwt malformed", http.StatusUnauthorized, "errors", data)
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseJSON)
			return
		}
		arrToken := strings.Split(h.IDToken, " ")
		if len(arrToken) != 2 {
			data := gin.H{"is_valid": false}
			responseJSON := helper.APIResponse("jwt malformed", http.StatusUnauthorized, "errors", data)
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseJSON)
			return
		}

		token, err := m.authService.ValidateToken(arrToken[1])
		if err != nil {
			data := gin.H{"is_valid": false}
			responseJSON := helper.APIResponse(err.Error(), http.StatusUnauthorized, "errors", data)
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseJSON)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			data := gin.H{"is_valid": false}
			responseJSON := helper.APIResponse("unauthorized", http.StatusUnauthorized, "errors", data)
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseJSON)
			return
		}
		userID := int(claims["user_id"].(float64))
		user, err := m.userService.FindUserById(userID)
		if err != nil {
			data := gin.H{"is_valid": false}
			responseJSON := helper.APIResponse(err.Error(), http.StatusUnauthorized, "errors", data)
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseJSON)
			return
		}
		c.Set("user", user)
	}
}
