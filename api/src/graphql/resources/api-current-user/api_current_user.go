package api_current_user

import (
	"github.com/benammann/drinkspot-core/api/utility"
	"github.com/gin-gonic/gin"
)

func Query_CurrentUser(ctx *gin.Context) (interface{}, error) {

	currentUser, err := utility.ExtractCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	return &CurrentUserResolver{
		CurrentUser{
			EmailAddress: &currentUser.EmailAddress,
		},
	}, nil

}
