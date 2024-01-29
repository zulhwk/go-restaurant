package domain

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	RestaurantDomain struct {
		ID           primitive.ObjectID `bson:"_id" json:"id"`
		Name         string             `json:"name"`
		RestaurantId string             `bson:"restaurant_id" json:"restaurant_id"`
		Cuisine      string             `json:"cuisine"`
		Address      interface{}        `json:"address"`
		Borough      string             `json:"borough"`
		Grades       []interface{}      `json:"grades"`
	}
	RestaurantCreateDomain struct {
		Name    string        `json:"name"`
		Cuisine string        `json:"cuisine"`
		Address interface{}   `json:"address"`
		Borough string        `json:"borough"`
		Grades  []interface{} `json:"grades"`
	}
	RestaurantCreateDomainResponse struct {
		ID      primitive.ObjectID `json:"id"`
		Name    string             `json:"name"`
		Cuisine string             `json:"cuisine"`
		Address interface{}        `json:"address"`
		Borough string             `json:"borough"`
		Grades  []interface{}      `json:"grades"`
	}
	RestaurantRepository interface {
		Create(ctx context.Context, payload *RestaurantCreateDomain) (primitive.ObjectID, error)
		Update()
		Delete()
		FindAll(ctx context.Context, filter primitive.M, opts *options.FindOptions) ([]RestaurantDomain, error)
		FindById(ctx context.Context, id primitive.ObjectID) RestaurantDomain
	}
	RestaurantUsecase interface {
		Create(ctx context.Context, payload *RestaurantCreateDomain) primitive.ObjectID
		Update()
		Delete()
		FindAll(ctx context.Context, opts *options.FindOptions) []RestaurantDomain
		FindById(ctx context.Context, id primitive.ObjectID) RestaurantDomain
	}
	RestaurantController interface {
		Create(c *fiber.Ctx) error
		Update()
		Delete()
		FindAll(c *fiber.Ctx) error
		FindById(c *fiber.Ctx) error
	}
)
