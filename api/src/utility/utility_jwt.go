package utility

import (
	"errors"
	"github.com/benammann/drinkspot-core/api/app/model"
	"github.com/benammann/drinkspot-core/api/database"
	"github.com/dgrijalva/jwt-go"
	"os"
)

// validates the given jwt token
func ValidateToken(token string) (*model.UserJWTClaims, bool, error) {
	// create the new claims struct
	claims := &model.UserJWTClaims{}

	// parse the claims
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// check for any errors
	if err != nil {
		return nil, false, err
	}

	// check if the token is valid
	if !tkn.Valid {
		return nil, false, nil
	}

	return claims, true, nil

}

// parses a new user using the given jwt user claims
func ParseUser(claims *model.UserJWTClaims) (*model.User, error) {

	// initialize the new user object
	requestedUser := &model.User{}

	// fetch the requested user
	if database.Current.Connection.Where("id = ?", claims.UserID).First(requestedUser).RecordNotFound() {
		return nil, errors.New("the requested user does not exist")
	}

	return requestedUser, nil

}
