package main

import (
	"log"
	"net/http"

	"github.com/cimartinez3/DEUNA"
	pb "github.com/cimartinez3/DEUNA/bank/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var grpcClient pb.BankServiceClient

func provideGrpcConnection() *grpc.ClientConn {
	conn, err := grpc.NewClient(DEUNA.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	grpcClient = pb.NewBankServiceClient(conn)

	return conn
}

func main() {
	conn := provideGrpcConnection()

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	http.HandleFunc("/charge", ChargeHandler)
	http.HandleFunc("/refund", RefundHandler)

	log.Println("API up in address: ", DEUNA.APIAddress)

	err := http.ListenAndServe(DEUNA.APIAddress, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
