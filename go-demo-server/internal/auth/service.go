package auth

import (
	"errors"
	"goadvancedserver/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepository *user.Repository
}

func (service *Service) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.FindByUsername(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	usr := user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}
	_, err = service.UserRepository.Create(&usr)
	if err != nil {
		return "", err
	}
	return usr.Email, nil
}

func (service *Service) Login(email, password string) (string, error) {
	usr, err := service.UserRepository.FindByUsername(email)
	if err != nil {
		return "", errors.New(ErrIncorrectEmailOrPassword)
	}
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrIncorrectEmailOrPassword)
	}
	return usr.Email, nil
}
