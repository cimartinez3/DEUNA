package main

import (
	"context"
	"github.com/cimartinez3/DEUNA/api/types"
	"github.com/cimartinez3/DEUNA/api/validator"
	pb "github.com/cimartinez3/DEUNA/bank/proto"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

// makeCharge makes some fields validations and send to gRPC server to do charge logic.
func makeCharge(c *gin.Context) {
	var charge types.ChargeRequest

	err := c.Bind(&charge)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.NewChargeValidator()
	if err = validate.ValidateTransaction(charge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := grpcClient.Charge(context.Background(), &pb.TransactionRequest{
		Id: uuid.Must(uuid.NewV4(), nil).String(),
		Card: &pb.Card{
			CardNumber:     charge.Card.Number,
			Cvv:            charge.Card.Cvv,
			ExpirationDate: charge.Card.ExpirationDate,
			CardHolder:     charge.Card.CardHolderName,
		},
		Merchant:   charge.Merchant,
		Amount:     charge.Amount,
		CustomerId: charge.Customer,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Println("CHARGE SUCCESSFULLY")

	c.JSON(http.StatusOK, response)
}

// getCharge get the charge_id from query param and send to gRPC to get the charge.
func getCharge(c *gin.Context) {
	id := c.Query("charge_id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id cant be empty"})
		return
	}

	response, err := grpcClient.GetCharge(context.Background(), &pb.ChargeId{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Println("FOUND CHARGE")

	c.JSON(http.StatusOK, response)
}
