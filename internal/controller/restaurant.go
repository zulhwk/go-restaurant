package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/zulhwk/go-restaurant/internal/domain"
	"github.com/zulhwk/go-restaurant/pkg/handlers"
	"github.com/zulhwk/go-restaurant/pkg/pagination"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RestaurantController struct {
	RestaurantUsecase domain.RestaurantUsecase
}

func NewRestaurantController(ru domain.RestaurantUsecase) *RestaurantController {
	return &RestaurantController{
		RestaurantUsecase: ru,
	}
}

func CreateRestaurantRoutes(app *fiber.App, rc *RestaurantController) {
	restaurant := app.Group("/restaurant")
	restaurant.Get("/", rc.FindAll)
}

func (rc *RestaurantController) Create() {}
func (rc *RestaurantController) Update() {}
func (rc *RestaurantController) Delete() {}
func (rc *RestaurantController) FindAll(c *fiber.Ctx) error {
	l, p := pagination.GetLimitAndPage(c.Query("limit"), c.Query("page"))
	fOpts := options.FindOptions{Limit: &l}
	data := rc.RestaurantUsecase.FindAll(context.Background(), &fOpts)
	handler := handlers.CreateWebResponse(handlers.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   data,
		Limit:  l,
		Page:   p,
	})
	return handler.WriteToResponseBody(c)
}
func (rc *RestaurantController) FindById() {}
