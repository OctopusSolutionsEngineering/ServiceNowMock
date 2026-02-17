package handler

import "github.com/gin-gonic/gin"

func View(c *gin.Context) {
	c.File("./html/changerequest.html")
}
