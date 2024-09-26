package main

import (
	"context"
	"testing"

	"github.com/cimartinez3/DEUNA/bank/proto"
	mocks "github.com/cimartinez3/DEUNA/mocks/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBankService_Refund(t *testing.T) {
	scenarios := []struct {
		name      string
		getTrxErr error
		refundErr error
		hasError  bool
	}{
		{
			name: "should make a refund successfully",
		}, {
			name:      "should return error when get transaction fails",
			getTrxErr: testError,
			hasError:  true,
		}, {
			name:      "should return error when make refund fails",
			refundErr: testError,
			hasError:  true,
		},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			assertions := assert.New(t)

			mongoSrv := &mocks.IMongo{}
			mongoSrv.On("GetTransaction", mock.Anything, mock.Anything, mock.Anything).Return(test.getTrxErr)
			mongoSrv.On("RefundAmounts", mock.Anything, mock.Anything).Return(test.refundErr)

			bs := BankService{mongoGtw: mongoSrv}

			res, err := bs.Refund(context.TODO(), &proto.RefundRequest{})

			if test.hasError {
				assertions.Error(err)
				assertions.Nil(res)
			} else {
				assertions.Nil(err)
			}
		})
	}
}
