package domain

import (
	"fmt"
	"github.com/nickdipreta/golang-microservices/mvc/utils"
	"net/http"
)

//assume that this is our database for now
var (
	users = map[int64]*User{
		123: &User{Id: 123, FirstName: "Nick", LastName: "DiPreta", Email: "nicholasdipreta2020@gmail.com"},
	}
)

//domain makes query against "database"
// if there is an error, we return the applciation error
func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf(" user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
