package main

import (
	"LibrApp/internal/db" // Переименовали импорт
	"LibrApp/internal/rest"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	db.InitDB()
	go startGRPC()

	r := mux.NewRouter()
	rest.RegisterRoutes(r)
	log.Println("REST server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func startGRPC() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpc.RegisterBookService(s)
	log.Println("gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
