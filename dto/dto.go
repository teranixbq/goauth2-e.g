package dto

import "goauth/model"

type ResponseLogin struct {
	Name string `json:"json"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ModelToResponse(data model.Users,token string)ResponseLogin{
	return ResponseLogin{
		Name: data.Name,
		Email: data.Email,
		Token: token,
	}
}

func ResponseToModel(data ResponseLogin) model.Users{
	return model.Users{
		Name: data.Name,
		Email: data.Email,
	}
}