package users

import (
	"net/http"
	"strconv"

	"github.com/aditya43/bookstore-oauth-go/oauth"
	"github.com/aditya43/golang-bookstore_users-api/domain/users"
	"github.com/aditya43/golang-bookstore_users-api/services"
	"github.com/aditya43/golang-bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.BadRequestErr("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, err := services.UserService.GetUser(userId)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if oauth.GetUserId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}

	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}

func Create(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestErr("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	res, restErr := services.UserService.CreateUser(user)

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, res.Marshall(oauth.IsPublic(c.Request)))
}

func Update(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequestErr("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestErr("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId
	isPatchMethod := c.Request.Method == http.MethodPatch
	res, restErr := services.UserService.UpdateUser(isPatchMethod, user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, res.Marshall(oauth.IsPublic(c.Request)))
}

func Delete(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequestErr("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	err := services.UserService.DeleteUser(userId)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	status := c.Query("status")

	users, err := services.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(oauth.IsPublic(c.Request)))
}

func Authenticate(c *gin.Context) {
	var loginRequest users.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		restErr := errors.BadRequestErr("Invalid credentials")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := services.UserService.Authenticate(&loginRequest)
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
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
