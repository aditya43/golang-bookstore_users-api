package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.String(http.StatusNotImplemented, "GET /users/:user_id")
}

func Create(c *gin.Context) {
	c.String(http.StatusNotImplemented, "POST /users")
}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "GET /users/search")
}
