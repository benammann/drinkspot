package controller

import (
	"github.com/benammann/drinkspot-core/api/app/model"
	"github.com/benammann/drinkspot-core/api/database"
	"github.com/benammann/drinkspot-core/api/utility"
	"github.com/benammann/drinkspot-core/api/validation"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type renewTokenRequest struct {
	Token string `json:"token"`
}

func Authentication(api *gin.Engine) {

	resource := api.Group("/auth")

	resource.POST("/login", func(ctx *gin.Context) {

		// parse the json request
		var json loginRequest
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var requestedUser model.User

		// fetch the requested user
		if database.Current.Connection.Where(&model.User{
			EmailAddress: json.Username,
		}).First(&requestedUser).RecordNotFound() {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Authentication Failed: Record not found"})
			return
		}

		jwtToken, err := requestedUser.Authenticate(json.Password)

		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"token": jwtToken,
		})
		return

	})

	// this endpoint allows to register a new account
	resource.POST("/register", func(ctx *gin.Context) {

		// parse the json request
		var json registerRequest
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validate the username
		if !validation.ValidateUsername(json.Username) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email Address"})
			return
		}

		// checks if the passwords match
		if !validation.ValidatePasswordConfirmation(json.Password, json.PasswordConfirm) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "The passwords do not match"})
			return
		}

		// checks if the password is strong enough
		if !validation.ValidatePasswordStrength(json.Password) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "The password is not strong enough"})
			return
		}

		var userCount int

		// check if the given user already exists
		database.Current.Connection.Model(&model.User{}).Where("email_address = ?", json.Username).Count(&userCount)

		if userCount == 0 {

			// initialize a new user
			newUser := &model.User{
				EmailAddress: json.Username,
			}

			// check if any errors are generated while encrypting the password
			err := newUser.SetPassword(json.Password)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
				return
			}

			// persist the user to the database
			database.Current.Connection.Create(newUser)

			// success
			ctx.JSON(200, gin.H{
				"error": nil,
			})

			return

		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
			return
		}

	})

	resource.POST("/renew-token", func(ctx *gin.Context) {

		// parse the json request
		var json renewTokenRequest
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validates the jwt token
		claims, valid, err := utility.ValidateToken(json.Token)

		// check for any errors
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "jwt signature invalid",
				})
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "could not parse jwt",
				})
			}
			return
		}

		// check if the token is valid
		if !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid JWT token",
			})
			return
		}

		// calculate the time until expiration
		tokenExpiresIn := time.Unix(claims.ExpiresAt, 0).Sub(time.Now())

		// allow to renew 30 seconds before expiration
		if tokenExpiresIn > 30*time.Second {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"token_expires_in": time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) / time.Second,
				"error":            "token renew is allowed 30 seconds before expiration",
			})
			return
		}

		// parse the current user based on the given request
		requestedUser, err := utility.ParseUser(claims)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// generate a new jwt token based on the current user
		newToken, err := requestedUser.GenerateJWTToken()

		// check for any errors
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate JWT token"})
			return
		}

		// return the new jwt token
		ctx.JSON(200, gin.H{
			"token": newToken,
		})

		return

	})

}
