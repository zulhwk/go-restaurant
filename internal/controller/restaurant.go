package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/zulhwk/go-restaurant/internal/domain"
	"github.com/zulhwk/go-restaurant/pkg/handlers"
	"github.com/zulhwk/go-restaurant/pkg/pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	restaurant.Post("/", rc.Create)
	restaurant.Get("/:id", rc.FindById)
}

/*
@Description Insert one restaurant.
@Params payload paramater domain.RestaurantCreateDomain
*/
func (rc *RestaurantController) Create(c *fiber.Ctx) error {
	requestBody := domain.RestaurantCreateDomain{}
	err := c.BodyParser(&requestBody)
	if err != nil {
		return nil
	}
	id := rc.RestaurantUsecase.Create(context.Background(), &requestBody)
	handler := handlers.CreateWebResponse(handlers.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success",
		Data: &domain.RestaurantCreateDomainResponse{
			ID:      id,
			Name:    requestBody.Name,
			Cuisine: requestBody.Cuisine,
			Address: requestBody.Address,
			Borough: requestBody.Borough,
			Grades:  requestBody.Grades,
		},
	})
	return handler.WriteToResponseBody(c)
}
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

func (rc *RestaurantController) FindById(c *fiber.Ctx) error {
	restaurantID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	data := rc.RestaurantUsecase.FindById(context.Background(), restaurantID)
	handler := handlers.CreateWebResponse(handlers.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   data,
	})
	return handler.WriteToResponseBody(c)
}
