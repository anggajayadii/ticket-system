package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleAuthorization allows only users with specific roles
func RoleAuthorization(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("user_role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not found"})
			return
		}

		userRole, ok := roleVal.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "role conversion failed"})
			return
		}

		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden access"})
	}
}
