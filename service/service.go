package service

import (
	"context"
	"errors"
	"goauth/model"
	"goauth/repository"
	"io"

	"golang.org/x/oauth2"
)

type service struct {
	repository repository.RepositoryInterface
	oauth      oauth2.Config
}

type ServiceInterface interface {
	CreateUser(data model.Users) error
	GetProfile(id string) (model.Users, error)
	GoogleAction() (string, error)
	GoogleCallback(code string) (string, error)
}

func NewService(repository repository.RepositoryInterface, oauth oauth2.Config) ServiceInterface {
	return &service{
		repository: repository,
	}
}

func (eg *service) CreateUser(data model.Users) error {

	err := eg.repository.CreateUser(data)
	if err != nil {
		return err
	}
	return nil

}

func (eg *service) GoogleAction() (string, error) {

	url := eg.oauth.AuthCodeURL("state")

	return url, nil
}

// func (eg *service) GoogleCallback(code string) (dto.ResponseLogin, error) {

// 	tokenOauth, err := eg.oauth.Exchange(context.Background(), code)
// 	if err != nil {
// 		return dto.ResponseLogin{}, errors.New("failed to exchange code for token: " + err.Error())
// 	}

// 	client := eg.oauth.Client(context.Background(), tokenOauth)
// 	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
// 	if err != nil {
// 		return dto.ResponseLogin{}, errors.New("failed to fetch user data: " + err.Error())
// 	}

// 	dataUsers, errCreate := eg.repository.CreateWithGoogle(resp.Body)
// 	if errCreate != nil {
// 		return dto.ResponseLogin{}, errCreate
// 	}

// 	JwtToken,errJwt := auth.CreateToken(dataUsers.Id)
// 	if errJwt != nil {
// 		return dto.ResponseLogin{},errJwt
// 	}

// 	response := dto.ModelToResponse(dataUsers, "")
// 	return response, nil
// }

func (eg *service) GoogleCallback(code string) (string, error) {

	tokenOauth, err := eg.oauth.Exchange(context.Background(), code)
	if err != nil {
		return "", errors.New("failed to exchange code for token: " + err.Error())
	}

	client := eg.oauth.Client(context.Background(), tokenOauth)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return "", errors.New("failed to fetch user data: " + err.Error())
	}
	defer resp.Body.Close()
	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("JSON Parsing Failed")
	}

	result := string(userData)
	return result, nil
}

func (eg *service) GetProfile(id string) (model.Users, error) {

	user, err := eg.repository.GetProfile(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
