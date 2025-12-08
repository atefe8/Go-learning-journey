package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gameproject/entity"
	"gameproject/pkg/phonenumber"
)

type repository interface {
	Register(u entity.User) (entity.User, error)
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
}

type Service struct {
	repository repository
}

func New(repo repository) Service {

	return Service{repository: repo}
}

// Register
type UserRegisterRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	User entity.User
}

func (s Service) Register(request UserRegisterRequest) (UserRegisterResponse, error) {

	if !phonenumber.IsVaild(request.Phone) {
		return UserRegisterResponse{}, fmt.Errorf("phone number is invaild 111")
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
	if len(request.Password) < 8 {
		return UserRegisterResponse{}, fmt.Errorf("password is invalid more than 3 charachter")

	}

	user := entity.User{
		Name:        request.Name,
		PhoneNumber: request.Phone,
		Password:    GetMD5Hash(request.Password),
	}

	//log.Fatalf("DEBUG USER: %+v", user)

	CreatedUser, error := s.repository.Register(user)
	if error != nil {
		return UserRegisterResponse{}, fmt.Errorf("there is error in create user", error.Error())
	}

	return UserRegisterResponse{CreatedUser}, nil

}

func GetMD5Hash(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

// login

type UserLoginRequest struct {
	PhoneNumber string `json:"phone"`
	Password    string `json:"password"`
}

type UserLoginResponse struct {
}

func (s Service) Login(req UserLoginRequest) (UserLoginResponse, error) {

	//exist , find user
	user, exist, error := s.repository.GetUserByPhoneNumber(req.PhoneNumber)

	if error != nil {
		return UserLoginResponse{}, fmt.Errorf("unexpected error %w", error)
	}

	if !exist {
		return UserLoginResponse{}, fmt.Errorf("user not found!")
	}

	if user.Password != GetMD5Hash(req.Password) {
		return UserLoginResponse{}, fmt.Errorf("user not found!")
	}

	return UserLoginResponse{}, nil
}
