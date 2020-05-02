package api_version

import (
	"github.com/gin-gonic/gin"
)

func Query_GetApiVersion(ctx *gin.Context) (interface{}, error) {

	appName := "drinkspot-api"
	appVersion := "v1.0.0"

	return &ApiVersionResolver{
		ApiVersion{
			AppName:    &appName,
			AppVersion: &appVersion,
		},
	}, nil

}
