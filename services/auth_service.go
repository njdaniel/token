package services

import (
	"encoding/json"
	"net/http"

	"github.com/njdaniel/token/services/models"
)

// Login authenicates user's username and password to return token if success
func Login(requestUser *models.User) (int, []byte) {
	authBackend := authenication.InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.UUID)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(parameters.TokenAuthenication{token})
			return http.StatusOk, response
		}
	}
	return http.StatusUnauthorized, []byte("")
}

// RefreshToken ...
func RefreshToken(requestUser *models.User) []byte {
	authBackend := authenication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.UUID)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthenication{token})
	if err != nil {
		panic(err)
	}
	return response
}

// Logout ...
func Logout(req *http.Request) error {
	authBackend := authenication.InitJWTAuthenticationBackend()
	tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}
