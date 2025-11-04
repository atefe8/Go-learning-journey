package userservice

import (
	"fmt"
	"game-application-new/GameProject/entity"
	"game-application-new/GameProject/pkg/phonenumber"
)

type repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	Repo repository
}
type RegisterRequest struct {
	Name  string
	Phone string
}

type RegisterResponse struct {
	User entity.User
}

func (s Service) Register(request RegisterRequest) (RegisterResponse, error) {

	if !phonenumber.IsVaild(request.Phone) {
		return RegisterResponse{}, fmt.Errorf("phone number is not vaild")
	}

	if isUnique, error := s.Repo.IsPhoneNumberUnique(request.Phone); error != nil || !isUnique {
		if error != nil {
			return RegisterResponse{}, fmt.Errorf("error happend", error)
		}
		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	if len(request.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name should be more than 3 charachter")

	}
	user := entity.User{
		Name:        request.Name,
		PhoneNumber: request.Phone,
	}

	createdUser, error := s.Repo.Register(user)

	if error != nil {
		return RegisterResponse{}, fmt.Errorf("there is an error in database, unexpected error", error)
	}

	return RegisterResponse{createdUser}, nil

}
