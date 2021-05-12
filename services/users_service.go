package services

import (
	"net/http"

	"github.com/cota-eng/go-users-api/domain/users"
	"github.com/cota-eng/go-users-api/utils/errors"
)


func CreateUser(user users.User) (*users.User,*errors.RestErr) {
	// no err and no user
	return nil,nil

	return &user,nil

	return &user,&errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}