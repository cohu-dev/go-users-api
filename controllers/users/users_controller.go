package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


var counter int

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"implement")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"implement")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"implement")
}