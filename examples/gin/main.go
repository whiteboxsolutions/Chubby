package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	chubby "github.com/whiteboxsolutions/Chubby"
	chubby_gin "github.com/whiteboxsolutions/Chubby/middleware/gin"
)

var rolls chubby.Rolls = chubby.New()

func main() {
	adminRoll := rolls.NewRoll("Admin")

	router := gin.Default()

	adminOnlyGroup := router.Group("/adminOnly")
	anyGroup := router.Group("/any")
	unauthorizedGroup := router.Group("/unauthorized")

	adminOnlyGroup.Use(AdminInjector())
	adminOnlyGroup.Use(chubby_gin.RollLimit(adminRoll))
	adminOnlyGroup.GET("/test", TestPath)

	anyGroup.GET("/test", TestPath)

	unauthorizedGroup.Use(chubby_gin.RollLimit(adminRoll))
	unauthorizedGroup.GET("/test", TestPath)

	router.Run("0.0.0.0:8090")
}

func TestPath(g *gin.Context) {
	g.JSON(http.StatusOK, "OK")
}

func AdminInjector() gin.HandlerFunc {
	return func(c *gin.Context) {
		r, err := rolls.Get("Admin")
		if err != nil {
			fmt.Println("Error getting admin roll")
			c.Abort()
		}
		c.Set("roll", r.Value)
	}
}
