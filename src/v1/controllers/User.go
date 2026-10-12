package v1

import (
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.String(200, "request get method:%s", c.Query("username"))
	case "POST":
		c.String(200, "request get method")
	default:
		c.String(200, "request method error")
	}
}

func ModifyUser(c *gin.Context) {
	c.String(200, "ModifyUser+")
}

func DelUser(c *gin.Context) {
	c.String(200, "DelUser+")
}
