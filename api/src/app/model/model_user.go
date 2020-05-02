package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"time"
)

type User struct {
	gorm.Model
	EmailAddress                  string `gorm:"unique_index"`
	EmailAddressConfirmed         bool
	EmailAddressConfirmationToken string
	PasswordHash                  string
	JWTRenewToken                 string
}

func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	u.JWTRenewToken = ""
	u.EmailAddressConfirmed = false
	u.GenerateEmailConfirmationToken()
	return nil
}

type UserJWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// Generates the confirmation token
func (u *User) GenerateEmailConfirmationToken() string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 35)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	token := string(b)
	u.EmailAddressConfirmationToken = token
	return token
}

// confirms the users email address
func (u *User) ConfirmEmailAddress(token string) error {
	if u.EmailAddressConfirmed {
		return errors.New("e-mail address already confirmed")
	} else if token == u.EmailAddressConfirmationToken {
		u.EmailAddressConfirmed = true
		u.EmailAddressConfirmationToken = ""
		return nil
	} else {
		return errors.New("invalid token")
	}
}

// validates the current password
func (u *User) ValidatePassword(currentPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(currentPassword)) == nil
}

// renews the password
func (u *User) RenewPassword(currentPassword string, newPassword string) error {

	// validates the password
	passwordValid := u.ValidatePassword(currentPassword)
	if !passwordValid {
		return errors.New("the current password is invalid")
	}

	// sets the new password
	return u.SetPassword(newPassword)

}

// sets the password
func (u *User) SetPassword(newPassword string) error {

	// generates a new password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	if err != nil {
		return err
	}

	// sets the password hash
	u.PasswordHash = string(hashedPassword)

	return nil

}

// Authenticates the given user and returns a jwt token and an error
func (u *User) Authenticate(currentPassword string) (string, error) {

	// validates the password hash against the plain password
	passwordValid := u.ValidatePassword(currentPassword)

	if !passwordValid {
		return "", errors.New("password invalid")
	}

	return u.GenerateJWTToken()

}

func (u *User) GenerateJWTToken() (string, error) {

	// creates the jwt token expiration time
	expirationTime := time.Now().Add(60 * time.Minute)

	// generates the token claims
	claims := &UserJWTClaims{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// generates the jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// encodes the token
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.New(err.Error())
	}

	return tokenString, nil

}
