package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cimartinez3/DEUNA"
	mongoGtw "github.com/cimartinez3/DEUNA/bank/mongo"
	pb "github.com/cimartinez3/DEUNA/bank/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)
import "go.mongodb.org/mongo-driver/mongo"

type BankService struct {
	pb.BankServiceServer
	mongoGtw mongoGtw.IMongo
}

const (
	dbName       = "DEUNA"
	merchants    = "merchants"
	customers    = "customers"
	transactions = "transactions"
	connString   = "mongodb://root:root@localhost:27017/"
)

var (
	merchantCollection     *mongo.Collection
	customerCollection     *mongo.Collection
	transactionsCollection *mongo.Collection
)

func main() {
	ctx := context.Background()
	InitMongoClient(ctx)
	InitializeData(ctx)

	lis, err := net.Listen("tcp", DEUNA.GrpcAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Listening on %s", DEUNA.GrpcAddress))

	s := grpc.NewServer()

	pb.RegisterBankServiceServer(s, &BankService{
		mongoGtw: mongoGtw.NewMongoGateway(merchantCollection, customerCollection, transactionsCollection),
	})

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}

func InitMongoClient(ctx context.Context) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}

	merchantCollection = client.Database(dbName).Collection(merchants)
	customerCollection = client.Database(dbName).Collection(customers)
	transactionsCollection = client.Database(dbName).Collection(transactions)
}

func InitializeData(ctx context.Context) {
	testCustomers := []interface{}{
		bson.D{{"customer_id", "CHARLIE1"},
			{"card", bson.D{{"cvv", "000"}, {"number", "40010000000002"}, {"expiration_date", "12/29"}, {"card_holder_name", "Carlos Martinez"}}},
			{"balance", 1000.0}},
	}

	_, err := customerCollection.InsertMany(ctx, testCustomers)
	if err != nil {
		log.Fatal(err)
		return
	}

	testMerchants := []interface{}{
		bson.D{{"name", "DEUNA"}, {"balance", 50000.0}},
	}

	_, err = merchantCollection.InsertMany(ctx, testMerchants)
	if err != nil {
		log.Fatal(err)
		return
	}
}
