package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Limit  int64       `json:"limit"`
	Page   int64       `json:"page"`
}

func CreateWebResponse(wr WebResponse) *WebResponse {
	return &WebResponse{
		Code:   wr.Code,
		Status: wr.Status,
		Data:   wr.Data,
		Limit:  wr.Limit,
		Page:   wr.Page,
	}
}

func (wr *WebResponse) WriteToResponseBody(c *fiber.Ctx) error {
	c.Accepts("application/json")
	return c.Status(wr.Code).JSON(wr)
}
