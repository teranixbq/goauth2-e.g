package handler

import (
	"goauth/model"
	"goauth/service"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.ServiceInterface
}

func Newhandler(service service.ServiceInterface) *handler {
	return &handler{service}
}

func (eg *handler) CreateUser(f *fiber.Ctx) error {
	input := model.Users{}
	if err := f.BodyParser(&input); err != nil {
		return f.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	err := eg.service.CreateUser(input)
	if err != nil {
		return f.Status(400).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}
	return f.Status(200).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (eg *handler) GoogleAction(f *fiber.Ctx) error {
	url, err := eg.service.GoogleAction()
	if err != nil {
		f.Status(400).JSON(fiber.Map{
			"message": "Failed to login",
		})
	}

	return f.Redirect(url, fiber.StatusFound)
}

func (eg *handler) GoogleCallback(f *fiber.Ctx) error {
	state := f.Query("state")
	if state != "state" {
		return f.SendString("States don't Match!!")
	}

	code := f.Query("code")

	result, err := eg.service.GoogleCallback(code)
	if err != nil {
		return f.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return f.Status(200).JSON(fiber.Map{
		"message": "succes",
		"data":    result,
	})

}
