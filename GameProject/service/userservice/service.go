package userservice

import (
	"fmt"
	"gameproject/entity"
	"gameproject/pkg/phonenumber"
)

type repository interface {
	Register(u entity.User) (entity.User, error)
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}
type Service struct {
	repository repository
}

type UserRegisterRequest struct {
	Name  string
	Phone string
}

type UserRegisterResponse struct {
	User entity.User
}

func (s Service) Register(request UserRegisterRequest) (UserRegisterResponse, error) {

	if !phonenumber.IsVaild(request.Phone) {
		return UserRegisterResponse{}, fmt.Errorf("phone number is invaild")
	}

	if isUnique, error := s.repository.IsPhoneNumberUnique(request.Phone); error != nil || !isUnique {
		if !isUnique {
			return UserRegisterResponse{}, fmt.Errorf("phone is not unique")
		}
		if error != nil {
			return UserRegisterResponse{}, fmt.Errorf("error happend")

		}
	}

	if len(request.Name) < 3 {
		return UserRegisterResponse{}, fmt.Errorf("name is invalid more than 3 charachter")

	}

	user := entity.User{
		Name:        request.Name,
		PhoneNumber: request.Phone,
	}

	CreatedUser, error := s.repository.Register(user)
	if error != nil {
		return UserRegisterResponse{}, fmt.Errorf("there is error in create user")
	}

	return UserRegisterResponse{CreatedUser}, nil

}
