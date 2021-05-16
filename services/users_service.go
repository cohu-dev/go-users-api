package services

import (
	"github.com/cota-eng/go-users-api/domain/users"
	"github.com/cota-eng/go-users-api/utils/errors"
)


func CreateUser(user users.User) (*users.User,*errors.RestErr) {
	if err:=user.Validate();err!=nil{
		return nil,err
	}
	if err:=user.Save();err!=nil{
		return nil,err
	}
	return &user,nil

	// return &user,nil

	// return &user,&errors.RestErr{
	// 	Status: http.StatusInternalServerError,
	// }
}