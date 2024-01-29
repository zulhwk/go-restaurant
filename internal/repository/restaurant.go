package repository

import (
	"context"

	"github.com/zulhwk/go-restaurant/internal/domain"
	"github.com/zulhwk/go-restaurant/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RestaurantRepository struct{}

func NewRestaurantRepository() domain.RestaurantRepository {
	return &RestaurantRepository{}
}

func (rr *RestaurantRepository) Create(ctx context.Context, payload *domain.RestaurantCreateDomain) (primitive.ObjectID, error) {
	collection, _ := mongodb.GetCollection("sample_restaurants", "restaurants")
	result, err := collection.InsertOne(ctx, payload)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (rr *RestaurantRepository) Update() {}
func (rr *RestaurantRepository) Delete() {}
func (rr *RestaurantRepository) FindAll(ctx context.Context, filter primitive.M, opts *options.FindOptions) ([]domain.RestaurantDomain, error) {
	collection, _ := mongodb.GetCollection("sample_restaurants", "restaurants")
	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return []domain.RestaurantDomain{}, err
	}
	results := []domain.RestaurantDomain{}
	for cur.Next(ctx) {
		var result domain.RestaurantDomain
		err := cur.Decode(&result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}
func (rr *RestaurantRepository) FindById() {}
