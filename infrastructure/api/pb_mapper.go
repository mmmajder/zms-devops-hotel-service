package api

import (
	"github.com/mmmajder/devops-booking-service/domain"
	pb "github.com/mmmajder/devops-booking-service/proto"
)

func mapHotel(hotel *domain.Hotel) *pb.Hotel {
	hotelPb := &pb.Hotel{
		Id:   hotel.Id.Hex(),
		Name: hotel.Name,
	}
	return hotelPb
}
