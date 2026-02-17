package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mcasperson/ServiceNowMock/internal/application/handler"
)

func main() {
	router := gin.Default()

	// Handler for /api/sn_chg_rest/change
	router.GET("/api/sn_chg_rest/change", handler.Change)

	// Handler for /api/sn_chg_rest/change/normal
	router.POST("/api/sn_chg_rest/change/normal", handler.Normal)

	// Handler for /api/sn_chg_rest/change/<any text>
	router.GET("/api/sn_chg_rest/change/:id", handler.ChangeId)

	router.PATCH("/api/sn_chg_rest/change/:id", handler.ChangeIdPatch)

	if err := router.Run(":8086"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
