package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OauthToken(c *gin.Context) {
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(`{
		"access_token": "blahblahblah",
		"refresh_token": "blahblahblah",
		"scope": "",
		"token_type": "Bearer",
		"expires_in": 1799
	}`))
}
