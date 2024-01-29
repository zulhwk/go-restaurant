package usecase

import (
	"context"

	"github.com/zulhwk/go-restaurant/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RestaurantUsecase struct {
	RestaurantRepository domain.RestaurantRepository
}

func NewRestaurantUsecase(rr domain.RestaurantRepository) domain.RestaurantUsecase {
	return &RestaurantUsecase{
		RestaurantRepository: rr,
	}
}

func (ru *RestaurantUsecase) Create() {}
func (ru *RestaurantUsecase) Update() {}
func (ru *RestaurantUsecase) Delete() {}
func (ru *RestaurantUsecase) FindAll(ctx context.Context, opts *options.FindOptions) []domain.RestaurantDomain {
	filter := bson.M{}
	data, _ := ru.RestaurantRepository.FindAll(ctx, filter, opts)
	return data
}
func (ru *RestaurantUsecase) FindById() {}
