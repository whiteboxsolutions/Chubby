package chubbyGin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	chubby "github.com/whiteboxsolutions/Chubby"
)

func RollLimit(requirement chubby.Roll) gin.HandlerFunc {
	return func(c *gin.Context) {
		roll, OK := c.Get("roll")
		if !OK {
			roll = 0
		}
		if chubby.HasRoll(roll.(uint), requirement) {
			// Pass on to the next-in-chain
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			c.Abort()
		}
	}
}
