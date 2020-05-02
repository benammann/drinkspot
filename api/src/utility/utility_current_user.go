package utility

import (
	"errors"
	"github.com/benammann/drinkspot-core/api/app/model"
	"github.com/gin-gonic/gin"
)

func ExtractCurrentUser(ctx *gin.Context) (*model.User, error) {

	// extract the current user
	value, exists := ctx.Get("current_user")

	// check if the user exists
	if !exists {
		return nil, errors.New("you are not allowed to use this method")
	}

	// cast the value as user
	return value.(*model.User), nil

}
