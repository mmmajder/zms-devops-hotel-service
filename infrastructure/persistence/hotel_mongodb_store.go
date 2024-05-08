package persistence

import (
	"context"
	"github.com/mmmajder/devops-booking-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "hoteldb"
	COLLECTION = "booking"
)

type HotelMongoDBStore struct {
	hotels *mongo.Collection
}

func NewHotelMongoDBStore(client *mongo.Client) domain.HotelStore {
	hotels := client.Database(DATABASE).Collection(COLLECTION)
	return &HotelMongoDBStore{
		hotels: hotels,
	}
}

func (store *HotelMongoDBStore) Get(id primitive.ObjectID) (*domain.Hotel, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *HotelMongoDBStore) GetAll() ([]*domain.Hotel, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *HotelMongoDBStore) Insert(hotel *domain.Hotel) error {
	hotel.Id = primitive.NewObjectID()
	result, err := store.hotels.InsertOne(context.TODO(), hotel)
	if err != nil {
		return err
	}
	hotel.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *HotelMongoDBStore) DeleteAll() {
	store.hotels.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *HotelMongoDBStore) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := store.hotels.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *HotelMongoDBStore) Update(id primitive.ObjectID, hotel *domain.Hotel) error {
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"name", hotel.Name},
		}},
	}
	_, err := store.hotels.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (store *HotelMongoDBStore) filter(filter interface{}) ([]*domain.Hotel, error) {
	cursor, err := store.hotels.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *HotelMongoDBStore) filterOne(filter interface{}) (hotel *domain.Hotel, err error) {
	result := store.hotels.FindOne(context.TODO(), filter)
	err = result.Decode(&hotel)
	return
}

func decode(cursor *mongo.Cursor) (hotels []*domain.Hotel, err error) {
	for cursor.Next(context.TODO()) {
		var hotel domain.Hotel
		err = cursor.Decode(&hotel)
		if err != nil {
			return
		}
		hotels = append(hotels, &hotel)
	}
	err = cursor.Err()
	return
}
