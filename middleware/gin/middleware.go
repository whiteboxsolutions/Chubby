package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	chubby "github.com/whiteboxsolutions/Chubby"
)

func RollLimit(c *gin.Context, roll chubby.Roll, requirement chubby.Roll) {
	if chubby.HasRoll(roll.Value, requirement) {
		// Pass on to the next-in-chain
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		c.Abort()
	}
}
