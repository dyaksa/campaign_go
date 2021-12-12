package middleware

import (
	"campaignproject/auth"
	"campaignproject/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

func AuthUser(auth auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var h authHeader
		err := c.ShouldBindHeader(&h)
		if err != nil {
			data := gin.H{"errors": err}
			responsJSON := helper.APIResponse("invalid arguments", http.StatusUnauthorized, "errors", data)
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

		_, err = auth.ValidateToken(arrToken[1])
		if err != nil {
			data := gin.H{"is_valid": false}
			responseJSON := helper.APIResponse(err.Error(), http.StatusUnauthorized, "errors", data)
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseJSON)
		}

	}
}
