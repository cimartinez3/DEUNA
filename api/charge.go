package main

import (
	"log"
	"net/http"

	"context"
	"encoding/json"
	"github.com/cimartinez3/DEUNA/api/types"
	"github.com/cimartinez3/DEUNA/api/validator"
	pb "github.com/cimartinez3/DEUNA/bank/proto"
	uuid "github.com/satori/go.uuid"
)

func ChargeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		makeCharge(w, r)
	case "GET":
		getCharge(w, r)
	default:
		http.Error(w, "wrong request", http.StatusBadRequest)
	}
}

func makeCharge(w http.ResponseWriter, r *http.Request) {
	var charge types.ChargeRequest

	err := json.NewDecoder(r.Body).Decode(&charge)
	if err != nil {
		http.Error(w, "wrong request", http.StatusBadRequest)
		return
	}

	validate := validator.NewChargeValidator()
	if err = validate.ValidateTransaction(charge); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{Message: err.Error()})

		return
	}

	log.Println("CHARGE SUCCESSFULLY")

	json.NewEncoder(w).Encode(response)
}

func getCharge(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("charge_id")

	if id == "" {
		http.Error(w, "charge id cant be empty", http.StatusBadRequest)
		return
	}

	response, err := grpcClient.GetCharge(context.Background(), &pb.ChargeId{Id: id})
	if err != nil {
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{Message: err.Error()})

		return
	}

	log.Println("FOUND CHARGE")

	json.NewEncoder(w).Encode(response)
}
