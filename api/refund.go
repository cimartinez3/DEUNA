package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	pb "github.com/cimartinez3/DEUNA/bank/proto"
)

// RefundHandler handles API petition and calls gRPC client to make a refund.
func RefundHandler(c *gin.Context) {
	id := c.Query("charge_id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id cant be empty"})
		return
	}

	_, err := grpcClient.Refund(context.Background(), &pb.RefundRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Println("REFUND SUCCESSFULLY")

	c.JSON(http.StatusOK, gin.H{"message": "REFUND SUCCESSFULLY"})
}
