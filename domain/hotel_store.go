package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelStore interface {
	Get(id primitive.ObjectID) (*Hotel, error)
	GetAll() ([]*Hotel, error)
	Insert(hotel *Hotel) error
	DeleteAll()
	Delete(id primitive.ObjectID) error
	Update(id primitive.ObjectID, hotel *Hotel) error
}
