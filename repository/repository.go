package repository

import (
	"encoding/json"
	"goauth/model"
	"io"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	GetProfile(id string) (model.Users, error)
	CreateWithGoogle(data io.Reader) (model.Users, error)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db: db,
	}
}

func (eg *repository) CreateUser(data model.Users) error {
	err := eg.db.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eg *repository) GetProfile(id string) (model.Users, error) {
	dataUsers := model.Users{}
	err := eg.db.Where("id = ?", id).Find(&dataUsers)
	if err.Error != nil {
		return model.Users{}, err.Error
	}
	return dataUsers, nil
}

func (eg *repository) CreateWithGoogle(data io.Reader) (model.Users, error) {
	dataUsers := model.Users{}
	dataGoogle := model.GoogleMethod{}

	err := json.NewDecoder(data).Decode(&dataGoogle)
	if err != nil {
		return model.Users{}, err
	}

	tx := eg.db.Where("email = ? ", dataGoogle.Email).Find(&dataUsers)
	if tx.RowsAffected == 1 {
		return dataUsers, tx.Error
	}

	if tx.Error != nil {
		return model.Users{}, tx.Error
	}

	// insert data to db
	input := model.GoogleToModel(dataGoogle)
	tx = eg.db.Create(&input)
	if tx.Error != nil {
		return model.Users{}, tx.Error
	}

	return input, nil
}
