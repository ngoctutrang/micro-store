package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ngoctutrang/micro-store/domain/users"
	"github.com/ngoctutrang/micro-store/services"
	"github.com/ngoctutrang/micro-store/utils/errors"
)

func GetUser(c *gin.Context) {
	userId, userIdErr := strconv.ParseInt(c.Params.ByName("user_id"), 10, 64)

	if userIdErr != nil {
		err := errors.NewBadRequestErr("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	if userId <= 0 {
		err := errors.NewBadRequestErr("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getUserErr := services.GetUser(userId)
	if getUserErr != nil {
		err := errors.NewBadRequestErr("Invalid user id")
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User

	if error := c.ShouldBindJSON(&user); error != nil {
		//TODO: handle sjon error
		restErr := errors.NewBadRequestErr("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	savedUser, savedUserErr := services.CreateUser(user)

	if savedUserErr != nil {
		//TODO: handle error
		c.JSON(savedUserErr.Status, savedUserErr)
		return
	}

	c.JSON(http.StatusCreated, savedUser)

}
