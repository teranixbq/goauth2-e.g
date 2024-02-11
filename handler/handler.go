package handler

import (
	"goauth/helper"
	"goauth/middleware"
	"goauth/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.ServiceInterface
}

func Newhandler(service service.ServiceInterface) *handler {
	return &handler{service}
}

func (eg *handler) GoogleAction(f *fiber.Ctx) error {
	url, err := eg.service.GoogleAction()
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			return f.Status(400).JSON(helper.ErrorResponse(err.Error()))
		}

		return f.Status(500).JSON(helper.ErrorResponse(err.Error()))
	}

	return f.Redirect(url, fiber.StatusFound)
}

func (eg *handler) GoogleCallback(f *fiber.Ctx) error {
	code := f.Query("code")

	result, err := eg.service.GoogleCallback(code)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			return f.Status(400).JSON(helper.ErrorResponse(err.Error()))
		}

		return f.Status(500).JSON(helper.ErrorResponse(err.Error()))
	}

	return f.Status(200).JSON(helper.SuccessWithDataResponse("succes",result))

}

func (eg *handler) GetProfileByID(f *fiber.Ctx) error {

	id, errExtract := middleware.ExtractToken(f)
	if errExtract != nil {
		return f.Status(400).JSON(helper.ErrorResponse(errExtract.Error()))
	}

	result, err := eg.service.GetProfile(id)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			return f.Status(400).JSON(helper.ErrorResponse(err.Error()))
		}
		return f.Status(500).JSON(helper.ErrorResponse(err.Error()))
	}

	return f.Status(200).JSON(helper.SuccessWithDataResponse("succes",result))
}
