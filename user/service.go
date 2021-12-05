package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterInput(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterInput(input RegisterUserInput) (User, error) {
	user := User{}
	user.Email = input.Email
	user.Name = input.Name
	user.Role = input.Role
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	existUser, err := s.repository.FindByEmail(email)
	if err != nil {
		return existUser, err
	}

	if existUser.ID == 0 {
		return existUser, errors.New("user not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.PasswordHash), []byte(password))
	if err != nil {
		return existUser, err
	}
	return existUser, nil
}
