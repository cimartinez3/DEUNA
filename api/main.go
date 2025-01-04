package main

import (
	"github.com/cimartinez3/DEUNA"
	pb "github.com/cimartinez3/DEUNA/bank/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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

	client := gin.Default()

	client.POST("/charge", func(c *gin.Context) {
		makeCharge(c)
	})

	client.GET("/charge", func(c *gin.Context) {
		getCharge(c)
	})

	client.POST("/refund", func(c *gin.Context) {
		RefundHandler(c)
	})

	log.Println("API up in address: ", DEUNA.APIAddress)

	err := client.Run(DEUNA.APIAddress)
	if err != nil {
		panic(err)
	}
}
