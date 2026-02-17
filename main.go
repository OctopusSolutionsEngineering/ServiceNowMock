package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Handler for /api/sn_chg_rest/change
	router.GET("/api/sn_chg_rest/change", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Change endpoint",
		})
	})

	// Handler for /api/sn_chg_rest/change/normal
	router.GET("/api/sn_chg_rest/change/normal", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Normal change endpoint",
		})
	})

	// Handler for /api/sn_chg_rest/change/<any text>
	router.GET("/api/sn_chg_rest/change/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Change endpoint with parameter",
			"id":      id,
		})
	})

	router.Run(":8080")
}
