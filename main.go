package main

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mcasperson/ServiceNowMock/internal/application/handler"
)

func getRequestBodyAsString(c *gin.Context) string {
	if c.Request == nil || c.Request.Body == nil {
		return ""
	}

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("failed to read request body: %v", err)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(nil))
		return ""
	}

	// Restore the body so downstream handlers can still bind/read it.
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return string(bodyBytes)
}

func logRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := getRequestBodyAsString(c)
		log.Printf("request method=%s path=%s query=%q body=%q", c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery, body)
		c.Next()
	}
}

func main() {
	router := gin.New()
	router.Use(logRequests())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", handler.Root)

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
