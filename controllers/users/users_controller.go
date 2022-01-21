package users

import (
	"github.com/gin-gonic/gin"
	"moviestore_users-api/domain/users"
	"moviestore_users-api/services"
	"moviestore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	//result, saveErr := services.UsersService.CreateUser(user)
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	//c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
	c.JSON(http.StatusCreated, result)
}
func Get(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("User Id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}
func Search(c *gin.Context) {

}
