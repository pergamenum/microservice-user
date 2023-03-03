package main

import (
	"log"
	"net"

	"github.com/pergamenum/microservice-user/pkg/payment"
	paymentpb "github.com/pergamenum/protobuf/golang/payment"
	"google.golang.org/grpc"
)

func main() {

	log.Println("Dingus Payment Microservice Starting")

	port := ":9000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	pc := payment.NewController()

	s := grpc.NewServer()
	paymentpb.RegisterCustomerServer(s, pc)
	paymentpb.RegisterOrderServer(s, pc)

	log.Println("Listening on port", port)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

}
