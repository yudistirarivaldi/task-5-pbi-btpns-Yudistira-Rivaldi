package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginInput) (User, error)
	IsEmailAvailable(email string) (User, error)
	GetUserByID(ID int) (User, error)
	GetAllUsers() ([]User, error)

}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	// mampping struct user ke struct user input

	user := User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash) //karena password hash awalnya byte jadi di pindah ke string
	user.Role = "user"

	newUser, err := s.repository.Save(user) //cek balikan dari repository save
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginInput) (User, error) {

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email) //cek balikannya dari repository findbyemail

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User tidak di temukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(email string) (User, error) {

	user, err := s.repository.FindByEmail(email) //cek balikannya dari repository findbyemail
	if err != nil {
		return user, err
	}

	if(user.Email != "") {
		return user, errors.New("User sudah ada")
	}

	return user, nil

}

func (s *service) GetUserByID(ID int) (User, error) {

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User tidak di temukan")
	}

	return user, nil

}

func (s *service) GetAllUsers() ([]User, error) {
	
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil

}