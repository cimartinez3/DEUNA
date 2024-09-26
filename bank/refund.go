package main

import (
	"context"
	"log"

	"github.com/cimartinez3/DEUNA/bank/proto"
	"github.com/cimartinez3/DEUNA/bank/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Refund gets balance of a charge and returns it to the customer.
func (b *BankService) Refund(ctx context.Context, request *proto.RefundRequest) (*emptypb.Empty, error) {
	charge := &types.Charge{}
	if err := b.mongoGtw.GetTransaction(ctx, request.Id, charge); err != nil {
		return nil, err
	}

	if err := b.mongoGtw.RefundAmounts(ctx, charge); err != nil {
		return nil, err
	}

	log.Println("REFUND SUCCESSFULLY")

	return &emptypb.Empty{}, nil

}
