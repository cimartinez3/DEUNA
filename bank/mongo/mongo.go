package mongo

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	"github.com/cimartinez3/DEUNA/bank/proto"
	"github.com/cimartinez3/DEUNA/bank/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMongo interface {
	CreateCustomer(ctx context.Context, req *proto.TransactionRequest) primitive.ObjectID
	CreateCharge(ctx context.Context, req *proto.TransactionRequest) primitive.ObjectID
	UpdateMerchantBalance(ctx context.Context, request *proto.TransactionRequest, merchant *types.Merchant) error
	UpdateCustomerBalance(ctx context.Context, request *proto.TransactionRequest, customer *types.CustomerItem) error
	GetCustomer(ctx context.Context, request *proto.TransactionRequest, customerItem *types.CustomerItem) error
	GetMerchant(ctx context.Context, request *proto.TransactionRequest, merchantItem *types.Merchant) error
	GetTransaction(ctx context.Context, id string, charge *types.Charge) error
	RefundAmounts(ctx context.Context, charge *types.Charge) error
}

type GatewayMongo struct {
	merchantCollection     *mongo.Collection
	customerCollection     *mongo.Collection
	transactionsCollection *mongo.Collection
}

func NewMongoGateway(m, c, t *mongo.Collection) IMongo {
	return &GatewayMongo{
		merchantCollection:     m,
		customerCollection:     c,
		transactionsCollection: t,
	}
}

func (g *GatewayMongo) CreateCustomer(ctx context.Context, req *proto.TransactionRequest) primitive.ObjectID {
	customer := types.CustomerItem{
		CustomerId: req.CustomerId,
		Card: types.CardItem{
			Cvv:            req.Card.Cvv,
			Number:         req.Card.CardNumber,
			ExpirationDate: req.Card.ExpirationDate,
			CardHolderName: req.Card.CardHolder,
		},
		Balance: 1000,
	}

	res, err := g.customerCollection.InsertOne(ctx, customer)
	if err != nil {
		log.Fatal(err)
		return primitive.ObjectID{}
	}

	oid, _ := res.InsertedID.(primitive.ObjectID)

	return oid
}

func (g *GatewayMongo) CreateCharge(ctx context.Context, req *proto.TransactionRequest) primitive.ObjectID {
	trx := types.Charge{
		From: req.CustomerId,
		Card: types.CardItem{
			Cvv:            req.Card.Cvv,
			Number:         req.Card.CardNumber,
			ExpirationDate: req.Card.ExpirationDate,
			CardHolderName: req.Card.CardHolder,
		},
		To:              req.Merchant,
		Amount:          req.Amount,
		TransactionType: "CHARGE",
	}

	res, err := g.transactionsCollection.InsertOne(ctx, trx)
	if err != nil {
		log.Fatal(err)
		return primitive.ObjectID{}
	}

	oid, _ := res.InsertedID.(primitive.ObjectID)

	return oid
}

func (g *GatewayMongo) UpdateMerchantBalance(ctx context.Context, request *proto.TransactionRequest, merchant *types.Merchant) error {
	update := bson.D{{"$inc", bson.D{{"balance", request.Amount}}}}

	_, err := g.merchantCollection.UpdateOne(ctx, bson.D{{"_id", merchant.Id}}, update)
	if err != nil {
		return err
	}
	return nil
}

func (g *GatewayMongo) UpdateCustomerBalance(ctx context.Context, request *proto.TransactionRequest, customer *types.CustomerItem) error {
	update := bson.D{{"$inc", bson.D{{"balance", request.Amount * -1}}}}

	_, err := g.customerCollection.UpdateOne(ctx, bson.D{{"_id", customer.Id}}, update)
	if err != nil {
		return err
	}

	return nil
}

func (g *GatewayMongo) GetCustomer(ctx context.Context, request *proto.TransactionRequest, customerItem *types.CustomerItem) error {
	return g.customerCollection.FindOne(ctx, bson.D{{"$and", []interface{}{
		bson.D{{"customer_id", request.CustomerId}},
		bson.D{{"card.number", request.Card.CardNumber}}}}}).Decode(customerItem)
}

func (g *GatewayMongo) GetMerchant(ctx context.Context, request *proto.TransactionRequest, merchantItem *types.Merchant) error {
	return g.merchantCollection.FindOne(ctx, bson.M{"name": request.Merchant}).Decode(merchantItem)
}

func (g *GatewayMongo) GetTransaction(ctx context.Context, id string, chargeItem *types.Charge) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return g.transactionsCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(chargeItem)
}

func (g *GatewayMongo) RefundAmounts(ctx context.Context, charge *types.Charge) error {
	query := bson.D{{"$and", []interface{}{
		bson.D{{"customer_id", charge.From}},
		bson.D{{"card.number", charge.Card.Number}}}}}

	merchantFilter := bson.D{{"$inc", bson.D{{"balance", charge.Amount}}}}

	_, err := g.customerCollection.UpdateOne(ctx, query, merchantFilter)
	if err != nil {
		return err
	}

	customerFilter := bson.D{{"$inc", bson.D{{"balance", charge.Amount * -1}}}}

	_, err = g.merchantCollection.UpdateOne(ctx, bson.D{{"name", charge.To}}, customerFilter)
	if err != nil {
		return err
	}

	charge.Id = primitive.ObjectID{}
	charge.TransactionType = "REFUND"

	_, err = g.transactionsCollection.InsertOne(ctx, charge)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
