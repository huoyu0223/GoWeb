package controllers

import (
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	c.String(200, "AddUser!")
}

func ModifyUser(c *gin.Context) {
	c.String(200, "ModifyUser!")
}

func DelUser(c *gin.Context) {
	c.String(200, "DelUser!")
}
