package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/cimartinez3/DEUNA/bank/proto"
	"github.com/cimartinez3/DEUNA/bank/types"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Charge maks a payment from a customer to a merchant.
func (b *BankService) Charge(ctx context.Context, request *proto.TransactionRequest) (*proto.TransactionResponse, error) {
	log.Println("Making new charge")

	customer := &types.CustomerItem{}
	merchant := &types.Merchant{}

	err := b.mongoGtw.GetCustomer(ctx, request, customer)
	if errors.Is(err, mongo.ErrNoDocuments) {
		customer.Id = b.mongoGtw.CreateCustomer(ctx, request)
	} else if err != nil {
		return handleErrorResponse(err)
	}

	if err = b.mongoGtw.GetMerchant(ctx, request, merchant); err != nil {
		return handleErrorResponse(err)
	}

	if err = b.mongoGtw.UpdateMerchantBalance(ctx, request, merchant); err != nil {
		return handleErrorResponse(err)
	}

	if err = b.mongoGtw.UpdateCustomerBalance(ctx, request, customer); err != nil {
		return handleErrorResponse(err)
	}

	id := b.mongoGtw.CreateCharge(ctx, request)
	log.Println("CHARGE SUCCESSFULLY")

	return &proto.TransactionResponse{Message: "SUCCESS", ChargeId: id.Hex()}, nil
}

// GetCharge returns specific charge by charge_id.
func (b *BankService) GetCharge(ctx context.Context, request *proto.ChargeId) (*proto.ChargeItem, error) {
	log.Println(fmt.Sprintf("Query charge %s", request.Id))
	charge := &types.Charge{}

	if err := b.mongoGtw.GetTransaction(ctx, request.Id, charge); err != nil {
		return nil, err
	}

	return &proto.ChargeItem{
		Id:   charge.Id.Hex(),
		From: charge.From,
		To:   charge.To,
		Card: &proto.Card{
			Cvv:            charge.Card.Cvv,
			CardNumber:     charge.Card.Number,
			ExpirationDate: charge.Card.ExpirationDate,
			CardHolder:     charge.Card.CardHolderName,
		},
		Amount: charge.Amount,
	}, nil

}

func handleErrorResponse(err error) (*proto.TransactionResponse, error) {
	log.Println("CHARGE DENIED DUE TO:")
	log.Println(err.Error())

	return &proto.TransactionResponse{Message: err.Error()}, status.Errorf(codes.Internal, err.Error())
}
