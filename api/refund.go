package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/cimartinez3/DEUNA/bank/proto"
)

func RefundHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("charge_id")

	if id == "" {
		http.Error(w, "charge id cant be empty", http.StatusBadRequest)
		return
	}

	_, err := grpcClient.Refund(context.Background(), &pb.RefundRequest{Id: id})
	if err != nil {
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{Message: err.Error()})

		return
	}

	log.Println("REFUND SUCCESSFULLY")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{Message: "SUCCESS REFUND"})
}
