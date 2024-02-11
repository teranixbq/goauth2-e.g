package service

import (
	"context"
	"errors"
	"goauth/dto"
	"goauth/middleware"
	"goauth/repository"

	"golang.org/x/oauth2"
)

type service struct {
	repository repository.RepositoryInterface
	oauth      oauth2.Config
}

type ServiceInterface interface {
	GoogleAction() (string, error)
	GoogleCallback(code string) (dto.ResponseLogin, error)
	GetProfile(id string) (dto.ResponseProfile, error)
}

func NewService(repository repository.RepositoryInterface, oauth oauth2.Config) ServiceInterface {
	return &service{
		repository: repository,
		oauth:      oauth,
	}
}

func (eg *service) GoogleAction() (string, error) {

	url := eg.oauth.AuthCodeURL("state")
	return url, nil
}

func (eg *service) GoogleCallback(code string) (dto.ResponseLogin, error) {

	tokenOauth, err := eg.oauth.Exchange(context.Background(), code)
	if err != nil {
		return dto.ResponseLogin{}, errors.New("failed to exchange code for token: " + err.Error())
	}

	client := eg.oauth.Client(context.Background(), tokenOauth)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return dto.ResponseLogin{}, errors.New(err.Error())
	}

	dataUsers, errCreate := eg.repository.CreateWithGoogle(resp.Body)
	if errCreate != nil {
		return dto.ResponseLogin{}, errCreate
	}

	JwtToken, errJwt := middleware.CreateToken(dataUsers.Id)
	if errJwt != nil {
		return dto.ResponseLogin{}, errJwt
	}

	response := dto.ModelToResponse(dataUsers, JwtToken)
	return response, nil
}

func (eg *service) GetProfile(id string) (dto.ResponseProfile, error) {
	dataUser, err := eg.repository.GetProfile(id)
	if err != nil {
		return dto.ResponseProfile{}, err
	}

	response := dto.ModelToResponseProfile(dataUser)
	return response, nil
}
