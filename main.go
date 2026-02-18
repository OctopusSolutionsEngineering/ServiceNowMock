package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mcasperson/ServiceNowMock/internal/application/handler"
)

func main() {
	router := gin.Default()

	router.POST("/oauth_token.do", handler.OauthToken)

	router.GET("change_request.do", handler.View)

	router.GET("/api/sn_chg_rest/change", handler.Change)

	router.POST("/api/sn_chg_rest/change/normal", handler.Normal)

	router.POST("/api/sn_chg_rest/change/emergency", handler.Emergency)

	router.GET("/api/sn_chg_rest/change/:id", handler.ChangeId)

	router.PATCH("/api/sn_chg_rest/change/:id", handler.ChangeIdPatch)

	if err := router.Run(":8086"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
