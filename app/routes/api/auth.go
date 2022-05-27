package api

import (
	"encoding/json"
	"golder/app/auth"
	"golder/app/consts"
	"golder/app/db"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthJSON struct {
	Username string
	Password string
}

func Auth(rw http.ResponseWriter, r *http.Request) {
	var data AuthJSON
	json.NewDecoder(r.Body).Decode(&data)

	var user db.User
	if res := consts.DB.First(&user, db.User{Username: data.Username}); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 403,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 403,
		})
		return
	}

	token, _ := auth.CreateAuthToken(user, time.Hour*24)

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"code":  200,
		"token": auth.SignAuthToken(token),
	})
}
