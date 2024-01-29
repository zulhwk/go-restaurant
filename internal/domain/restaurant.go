package domain

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	RestaurantDomain struct {
		ID           primitive.ObjectID `bson:"_id"`
		Name         string
		RestaurantId string `bson:"restaurant_id"`
		Cuisine      string
		Address      interface{}
		Borough      string
		Grades       []interface{}
	}
	RestaurantRepository interface {
		Create()
		Update()
		Delete()
		FindAll(ctx context.Context, filter primitive.M, opts *options.FindOptions) ([]RestaurantDomain, error)
		FindById()
	}
	RestaurantUsecase interface {
		Create()
		Update()
		Delete()
		FindAll(ctx context.Context, opts *options.FindOptions) []RestaurantDomain
		FindById()
	}
	RestaurantController interface {
		Create()
		Update()
		Delete()
		FindAll(c *fiber.Ctx) error
		FindById()
	}
)
