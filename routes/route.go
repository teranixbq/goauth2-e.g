package routes

import (
	"goauth/handler"
	"goauth/repository"
	"goauth/service"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func RouteInit(f *fiber.App, db *gorm.DB, oauth oauth2.Config) {
	userRepository := repository.NewRepository(db)
	userService := service.NewService(userRepository,oauth)
	userHandler := handler.Newhandler(userService)

	f.Post("/register", userHandler.CreateUser)
	//f.Get("/profile/:id", userHandler.GetProfile)
	f.Get("/google_login", userHandler.GoogleAction)
	f.Get("/google_callback", userHandler.GoogleCallback)
}	