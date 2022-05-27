package auth

import (
	"errors"
	"golder/app/consts"
	"golder/app/db"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthClaim struct {
	ID       uint
	Username string
	Auth     string
	jwt.StandardClaims
}

func SignAuthToken(token *jwt.Token) string {
	tString, _ := token.SignedString([]byte(consts.Config.Authentication.SECRET_KEY))
	return tString
}

func CreateAuthToken(user db.User, validTime time.Duration) (*jwt.Token, error) {
	claim := AuthClaim{
		ID:       user.ID,
		Username: user.Username,
		Auth:     user.Auth,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(validTime).Unix(),
			Issuer:    consts.App.Name + "-" + consts.App.Version,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim), nil
}

func VerifyAuthToken(tokenString string) (*AuthClaim, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&AuthClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(consts.Config.Authentication.SECRET_KEY), nil
		},
	)
	if err != nil {
		return &AuthClaim{}, err
	}

	claim, _ := token.Claims.(*AuthClaim)
	if claim.ExpiresAt < time.Now().Unix() {
		return &AuthClaim{}, errors.New("token expired")
	}

	return claim, nil
}

func VerifyAuthHeader(r *http.Request) (*AuthClaim, error) {
	return VerifyAuthToken(strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
}

func VerifyAuthCookie(r *http.Request) (*AuthClaim, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}

	return VerifyAuthToken(cookie.Value)
}
