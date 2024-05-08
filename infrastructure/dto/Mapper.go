package dto

import (
	"github.com/mmmajder/zms-devops-hotel-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapHotel(id primitive.ObjectID, hotel *HotelDto) *domain.Hotel {
	hotelPb := &domain.Hotel{
		Id:   id,
		Name: hotel.Name,
	}
	return hotelPb
}
