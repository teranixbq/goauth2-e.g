package routes

import (
	"goauth/handler"
	"goauth/repository"
	"goauth/service"
	"goauth/middleware"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func RouteInit(f *fiber.App, db *gorm.DB, oauth oauth2.Config) {
	userRepository := repository.NewRepository(db)
	userService := service.NewService(userRepository,oauth)
	userHandler := handler.Newhandler(userService)

	f.Get("/auth/google", userHandler.GoogleAction)
	f.Get("/callback", userHandler.GoogleCallback)
	f.Get("/profile",middleware.JWTMiddleware(),userHandler.GetProfileByID)
}	