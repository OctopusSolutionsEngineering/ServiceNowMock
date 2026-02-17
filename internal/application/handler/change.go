package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Change(c *gin.Context) {
	query := c.Query("sysparm_query")

	// JSON escape the query string
	escapedQuery, _ := json.Marshal(query)
	// Remove the surrounding quotes added by json.Marshal
	escapedQueryStr := string(escapedQuery[1 : len(escapedQuery)-1])

	c.JSON(http.StatusOK, `{
		"result": [
			{
				"__meta": {
					"encodedQuery": "short_description=`+escapedQueryStr+`",
					"fields": {
						"applied": [
							"short_description"
						],
						"ignored": []
					}
				}
			}
		]
	}`)
}
