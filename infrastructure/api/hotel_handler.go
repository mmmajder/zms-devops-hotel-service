package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mmmajder/devops-booking-service/application"
	"github.com/mmmajder/devops-booking-service/domain"
	"github.com/mmmajder/devops-booking-service/infrastructure/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type HotelHandler struct {
	service *application.HotelService
}

type HotelsResponse struct {
	Hotels []*domain.Hotel `json:"hotels"`
	Size   string          `json:"size"`
}

type HotelResponse struct {
	Hotel *domain.Hotel `json:"hotel"`
}

func NewHotelHandler(service *application.HotelService, mux *runtime.ServeMux) *HotelHandler {
	server := &HotelHandler{
		service: service,
	}
	return server
}

func (handler *HotelHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/booking/hotels", handler.GetDetails)
	err = mux.HandlePath("GET", "/booking/hotels/{id}", handler.GetById)
	err = mux.HandlePath("POST", "/booking/hotels", handler.AddHotel)
	err = mux.HandlePath("PUT", "/booking/hotels/{id}", handler.UpdateHotel)
	err = mux.HandlePath("DELETE", "/booking/hotels/{id}", handler.DeleteHotel)
	err = mux.HandlePath("GET", "/booking", handler.GetHealthCheck)
	if err != nil {
		panic(err)
	}
}

func (handler *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	hotelId, err := primitive.ObjectIDFromHex(pathParams["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedHotelDto dto.HotelDto
	err = json.NewDecoder(r.Body).Decode(&updatedHotelDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(updatedHotelDto)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	updatedHotel := dto.MapHotel(hotelId, &updatedHotelDto)

	err = handler.service.Update(hotelId, updatedHotel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *HotelHandler) DeleteHotel(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	hotelId, err := primitive.ObjectIDFromHex(pathParams["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.service.Delete(hotelId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *HotelHandler) GetHealthCheck(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	response := HotelsResponse{
		Size: "HOTEL SERVICE OK",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (handler *HotelHandler) GetById(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	hotelId, err := primitive.ObjectIDFromHex(pathParams["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hotel, err := handler.service.Get(hotelId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response := HotelResponse{
		Hotel: hotel,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (handler *HotelHandler) GetDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	hotel, err := handler.service.GetAll()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response := HotelsResponse{
		Hotels: hotel,
		Size:   "HOTEL SERVICE",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (handler *HotelHandler) AddHotel(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var newHotel domain.Hotel
	err := json.NewDecoder(r.Body).Decode(&newHotel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.service.Add(&newHotel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
