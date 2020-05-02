package middleware

import (
	"github.com/benammann/drinkspot-core/api/utility"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthenticateUser(ctx *gin.Context) {

	reqToken := ctx.GetHeader("Authorization")

	if reqToken != "" {

		splitToken := strings.Split(reqToken, "Bearer")

		if len(splitToken) != 2 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid authorization header",
			})
			ctx.Abort()
			return
		}

		jwtToken := strings.TrimSpace(splitToken[1])

		// validates the jwt token
		claims, valid, err := utility.ValidateToken(jwtToken)

		// check for any errors
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "jwt signature invalid",
				})
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"token": jwtToken,
					"error": err.Error(),
				})
			}
			ctx.Abort()
			return
		}

		// check if the token is valid
		if !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid JWT token",
			})
			ctx.Abort()
			return
		}

		currentUser, err := utility.ParseUser(claims)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not parse current user",
			})
			ctx.Abort()
			return
		}

		// set the current user
		ctx.Set("current_user", currentUser)

	}

	ctx.Next()

}
