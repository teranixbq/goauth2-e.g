package dto

import "goauth/model"

type ResponseLogin struct {
	Name  string `json:"json"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type ResponseProfile struct {
	Id    string `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ModelToResponse(data model.Users, token string) ResponseLogin {
	return ResponseLogin{
		Name:  data.Name,
		Email: data.Email,
		Token: token,
	}
}

func ModelToResponseProfile(data model.Users) ResponseProfile {
	return ResponseProfile{
		Id:    data.Id,
		Image: data.Image,
		Name:  data.Name,
		Email: data.Email,
	}
}
