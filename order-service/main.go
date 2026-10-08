package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/Mizuhar4/order-service/pb"
)

type server struct {
	pb.UnimplementedOrderServiceServer
}

func (s *server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return &pb.GetOrderResponse{
		Order: &pb.Order{
			OrderId:       req.OrderId,
			ClientId:      "client-uuid-123",
			RestaurantId:  "rest-uuid-456",
			Status:        "CREATED",
			DeliveryAddress: "Av. Universidad 01230",
			TotalAmount:   14990,
			CreatedAt:     "2026-10-07T20:00:00Z",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("no se pudo escuchar: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterOrderServiceServer(srv, &server{})

	log.Println("Servidor gRPC de Order Service escuchando en :50051")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error al servir: %v", err)
	}
}
