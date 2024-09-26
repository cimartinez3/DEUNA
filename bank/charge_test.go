package main

import (
	"context"
	"errors"
	"testing"

	"github.com/cimartinez3/DEUNA/bank/proto"
	mocks "github.com/cimartinez3/DEUNA/mocks/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var testError = errors.New("this is a fake error")

func TestBankService_Charge(t *testing.T) {
	scenarios := []struct {
		name                     string
		getCustomerErr           error
		getMerchantErr           error
		updateMerchantBalanceErr error
		updateCustomerBalanceErr error
		hasError                 bool
	}{
		{
			name: "should run charge logic successfully",
		}, {
			name:           "should create customer and continue logic when query brings empty",
			getCustomerErr: mongo.ErrNoDocuments,
		}, {
			name:           "should return error when query customer fails",
			getCustomerErr: testError,
			hasError:       true,
		}, {
			name:           "should return error when query merchant fails",
			getMerchantErr: testError,
			hasError:       true,
		}, {
			name:                     "should return error when update merchant  balance fails",
			updateMerchantBalanceErr: testError,
			hasError:                 true,
		}, {
			name:                     "should return error when update customer  balance fails",
			updateCustomerBalanceErr: testError,
			hasError:                 true,
		},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			assertions := assert.New(t)

			mongoSrv := &mocks.IMongo{}

			mongoSrv.On("GetCustomer", mock.Anything, mock.Anything, mock.Anything).
				Return(test.getCustomerErr)
			mongoSrv.On("CreateCustomer", mock.Anything, mock.Anything).
				Return(primitive.NewObjectID())
			mongoSrv.On("GetMerchant", mock.Anything, mock.Anything, mock.Anything).
				Return(test.getMerchantErr)
			mongoSrv.On("UpdateMerchantBalance", mock.Anything, mock.Anything, mock.Anything).
				Return(test.updateMerchantBalanceErr)
			mongoSrv.On("UpdateCustomerBalance", mock.Anything, mock.Anything, mock.Anything).
				Return(test.updateCustomerBalanceErr)
			mongoSrv.On("CreateCharge", mock.Anything, mock.Anything).
				Return(primitive.NewObjectID())

			bs := BankService{mongoGtw: mongoSrv}

			res, err := bs.Charge(context.TODO(), &proto.TransactionRequest{})

			if test.hasError {
				assertions.Error(err)
			} else {
				assertions.Equal("SUCCESS", res.Message)
			}
		})
	}
}

func TestBankService_GetCharge(t *testing.T) {
	scenarios := []struct {
		name      string
		getTrxErr error
		hasError  bool
	}{
		{
			name: "should return a charge successfully",
		}, {
			name:      "should return error when get transaction fails",
			getTrxErr: testError,
			hasError:  true,
		},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			assertions := assert.New(t)

			mongoSrv := &mocks.IMongo{}
			mongoSrv.On("GetTransaction", mock.Anything, mock.Anything, mock.Anything).Return(test.getTrxErr)

			bs := BankService{mongoGtw: mongoSrv}

			res, err := bs.GetCharge(context.TODO(), &proto.ChargeId{})

			if test.hasError {
				assertions.Nil(res)
				assertions.Error(err)
			} else {
				assertions.Nil(err)
				assertions.IsType(&proto.ChargeItem{}, res)
			}
		})
	}
}
