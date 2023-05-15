package controllers

import (
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.String(200, "request get method:%s", c.Query("username"))
	} else if c.Request.Method == "POST" {
		c.String(200, "request get method")
	} else {
		c.String(200, "request method error")
	}
}

func ModifyUser(c *gin.Context) {
	c.String(200, "ModifyUser+")
}

func DelUser(c *gin.Context) {
	c.String(200, "DelUser+")
}
