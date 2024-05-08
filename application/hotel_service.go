package application

import (
	"github.com/mmmajder/devops-booking-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelService struct {
	store domain.HotelStore
}

func NewHotelService(store domain.HotelStore) *HotelService {
	return &HotelService{
		store: store,
	}
}

func (service *HotelService) Get(id primitive.ObjectID) (*domain.Hotel, error) {
	return service.store.Get(id)
}

func (service *HotelService) GetAll() ([]*domain.Hotel, error) {
	return service.store.GetAll()
}

func (service *HotelService) Add(hotel *domain.Hotel) error {
	err := service.store.Insert(hotel)
	if err != nil {
		return err
	}
	return nil
}

func (service *HotelService) Update(id primitive.ObjectID, hotel *domain.Hotel) error {
	_, err := service.store.Get(id)
	if err != nil {
		return err // Return error if hotel does not exist
	}
	err = service.store.Update(id, hotel)
	if err != nil {
		return err
	}

	return nil
}

func (service *HotelService) Delete(id primitive.ObjectID) error {
	_, err := service.store.Get(id)
	if err != nil {
		return err // Return error if hotel does not exist
	}
	err = service.store.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
