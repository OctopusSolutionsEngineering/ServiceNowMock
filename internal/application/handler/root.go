package handler

import "github.com/gin-gonic/gin"

func Root(c *gin.Context) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	fullURL := scheme + "://" + c.Request.Host + c.Request.RequestURI
	c.String(200, "Configure your Octopus ServiceNow connection with "+fullURL+" to mock the interaction with a ServiceNow instance.")
}
