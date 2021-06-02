// MVC Driven Development

package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hammaad90/bookstore_users-api/domain/users"
	"github.com/hammaad90/bookstore_users-api/services"
	"github.com/hammaad90/bookstore_users-api/utils/errors"
)

// every controller that we have when using gin-gonic http framework needs to have this  c *gin.Context interface

func getUserId(userIdParam string) (int64, *errors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Get(c *gin.Context) {
	// c.String(http.StatusNotImplemented, format:"implement me!")
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusCreated, user.Marshall(c.GetHeader("x-Public") == "true"))

}

func Create(c *gin.Context) {
	var user users.User
	fmt.Println("rrrrrrrrrrrrrrrrrrr", user)
	// handling req.body error if json is not accurate
	// this is reading req.body and handling json error if any in req/.body json
	if err := c.ShouldBindJSON(&user); err != nil {
		// handling json errors
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// if no error sending user to service to get created
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {

		// if error in processing user
		c.JSON(saveErr.Status, saveErr)
		return
	}

	// if user is created successfully
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("x-Public") == "true"))

}

// Update user
func Update(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// handling json errors
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("x-Public") == "true"))

}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})

}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("x-Public") == "true"))
}
