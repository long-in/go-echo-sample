package services

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
}

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		jsonFile, err := os.Open("./config/users.json")
		if err != nil {
			return false, err
		}
		defer jsonFile.Close()

		var users Users
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &users)
		bytePass := sha512.Sum512([]byte(password))
		password = fmt.Sprintf("%x", bytePass)
		for i := 0; i < len(users.Users); i++ {
			userName, userPassword := users.Users[i].Name, users.Users[i].Password
			if userName == username && userPassword == password {
				return true, nil
			}
		}

		return false, nil

	})
}
