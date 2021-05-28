package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hammaad90/bookstore_users-api/domain/users"
	"github.com/hammaad90/bookstore_users-api/services"
	"github.com/hammaad90/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	// c.String(http.StatusNotImplemented, format:"implement me!")
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		//TODO: handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)

}

func CreateUser(c *gin.Context) {
	var user users.User
	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO: Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	// TODO:  handle json error
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO: return bad request to caller
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)

}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, format:"implement me!")

// }
