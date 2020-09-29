package services

import (
	"github.com/nickdipreta/golang-microservices/mvc/domain"
	"github.com/nickdipreta/golang-microservices/mvc/utils"
)
// service is calling the domain
func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
