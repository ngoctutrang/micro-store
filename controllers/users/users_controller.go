package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctutrang/micro-store/domain/users"
	"github.com/ngoctutrang/micro-store/errors"
	"github.com/ngoctutrang/micro-store/services"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}

func CreateUser(c *gin.Context) {
	var user users.User

	if error := c.ShouldBindJSON(&user); error != nil {
		//TODO: handle sjon error
		restErr := errors.RestEff{
			Message: "Invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}
	createdUser, createdUserErr := services.CreateUser(user)

	if createdUserErr != nil {
		//TODO: handle error
		return
	}

	c.JSON(http.StatusCreated, createdUser)

}
