package model

type Users struct {
	Id       string `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Password string `json:"password"`
}

type GoogleMethod struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"picture"`
}

func GoogleToModel(data GoogleMethod) Users {
	return Users{
		Id:    data.Id,
		Name:  data.Name,
		Email: data.Email,
		Image: data.Image,
	}
}
