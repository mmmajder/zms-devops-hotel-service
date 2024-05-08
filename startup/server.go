package startup

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mmmajder/zms-devops-hotel-service/application"
	"github.com/mmmajder/zms-devops-hotel-service/domain"
	"github.com/mmmajder/zms-devops-hotel-service/infrastructure/api"
	"github.com/mmmajder/zms-devops-hotel-service/infrastructure/persistence"
	booking "github.com/mmmajder/zms-devops-hotel-service/proto"
	"github.com/mmmajder/zms-devops-hotel-service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type Server struct {
	config *config.Config
	mux    *runtime.ServeMux
}

func NewServer(config *config.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	return server
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	hotelStore := server.initHotelStore(mongoClient)
	hotelService := server.initHotelService(hotelStore)
	hotelHandler := server.initHotelHandler(hotelService)
	hotelHandler.Init(server.mux)
	bookingHandler := server.initBookingHandler(hotelService)
	go server.startGrpcServer(bookingHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}

func (server *Server) startGrpcServer(bookingHandler *api.BookingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	booking.RegisterBookingServiceServer(grpcServer, bookingHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.HotelDBUsername, server.config.HotelDBPassword, server.config.HotelDBHost, server.config.HotelDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initHotelStore(client *mongo.Client) domain.HotelStore {
	store := persistence.NewHotelMongoDBStore(client)
	store.DeleteAll()
	for _, hotel := range hotels {
		err := store.Insert(hotel)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initHotelService(store domain.HotelStore) *application.HotelService {
	return application.NewHotelService(store)
}

func (server *Server) initHotelHandler(service *application.HotelService) *api.HotelHandler {
	return api.NewHotelHandler(service, server.mux)
}

func (server *Server) initBookingHandler(service *application.HotelService) *api.BookingHandler {
	return api.NewBookingHandler(service)
}
