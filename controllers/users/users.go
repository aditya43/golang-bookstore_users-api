package users

import (
	"net/http"

	"github.com/aditya43/golang-bookstore_users-api/domain/users"
	"github.com/aditya43/golang-bookstore_users-api/services"
	"github.com/aditya43/golang-bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.String(http.StatusNotImplemented, "GET /users/:user_id")
}

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestErr("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	res, restErr := services.CreateUser(user)

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "GET /users/search")
}

/*
// One way to unmarshall JSON and populate user struct with the values
func Create(c *gin.Context) {
	var user users.User

	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		// TODO: Handle error
		fmt.Println(err.Error())
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		// TODO: Handle json error
		fmt.Println(err.Error())
		return
	}
	res, err := services.CreateUser(user)

	if err != nil {
		// TODO: Handle user creation error
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}
*/
