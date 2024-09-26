// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: bank/proto/bank.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardNumber     string `protobuf:"bytes,1,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	Cvv            string `protobuf:"bytes,2,opt,name=cvv,proto3" json:"cvv,omitempty"`
	ExpirationDate string `protobuf:"bytes,3,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	CardHolder     string `protobuf:"bytes,4,opt,name=card_holder,json=cardHolder,proto3" json:"card_holder,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_bank_proto_bank_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *Card) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

func (x *Card) GetExpirationDate() string {
	if x != nil {
		return x.ExpirationDate
	}
	return ""
}

func (x *Card) GetCardHolder() string {
	if x != nil {
		return x.CardHolder
	}
	return ""
}

type ChargeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	From   string  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To     string  `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Card   *Card   `protobuf:"bytes,4,opt,name=card,proto3" json:"card,omitempty"`
	Amount float32 `protobuf:"fixed32,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ChargeItem) Reset() {
	*x = ChargeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeItem) ProtoMessage() {}

func (x *ChargeItem) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeItem.ProtoReflect.Descriptor instead.
func (*ChargeItem) Descriptor() ([]byte, []int) {
	return file_bank_proto_bank_proto_rawDescGZIP(), []int{1}
}

func (x *ChargeItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChargeItem) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *ChargeItem) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ChargeItem) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *ChargeItem) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Card       *Card   `protobuf:"bytes,2,opt,name=card,proto3" json:"card,omitempty"`
	Merchant   string  `protobuf:"bytes,3,opt,name=merchant,proto3" json:"merchant,omitempty"`
	Amount     float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	CustomerId string  `protobuf:"bytes,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_bank_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TransactionRequest) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *TransactionRequest) GetMerchant() string {
	if x != nil {
		return x.Merchant
	}
	return ""
}

func (x *TransactionRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type ChargeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChargeId) Reset() {
	*x = ChargeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeId) ProtoMessage() {}

func (x *ChargeId) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_bank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeId.ProtoReflect.Descriptor instead.
func (*ChargeId) Descriptor() ([]byte, []int) {
	return file_bank_proto_bank_proto_rawDescGZIP(), []int{3}
}

func (x *ChargeId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RefundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RefundRequest) Reset() {
	*x = RefundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundRequest) ProtoMessage() {}

func (x *RefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundRequest.ProtoReflect.Descriptor instead.
func (*RefundRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_bank_proto_rawDescGZIP(), []int{4}
}

func (x *RefundRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ChargeId string `protobuf:"bytes,2,opt,name=charge_id,json=chargeId,proto3" json:"charge_id,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_bank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_bank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_bank_proto_bank_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TransactionResponse) GetChargeId() string {
	if x != nil {
		return x.ChargeId
	}
	return ""
}

var File_bank_proto_bank_proto protoreflect.FileDescriptor

var file_bank_proto_bank_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x04, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x63, 0x76, 0x76, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x22, 0x78, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61,
	0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x12, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49,
	0x64, 0x32, 0xb2, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x12, 0x0e, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x65, 0x7a, 0x33,
	0x2f, 0x44, 0x45, 0x55, 0x4e, 0x41, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_proto_bank_proto_rawDescOnce sync.Once
	file_bank_proto_bank_proto_rawDescData = file_bank_proto_bank_proto_rawDesc
)

func file_bank_proto_bank_proto_rawDescGZIP() []byte {
	file_bank_proto_bank_proto_rawDescOnce.Do(func() {
		file_bank_proto_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_proto_bank_proto_rawDescData)
	})
	return file_bank_proto_bank_proto_rawDescData
}

var file_bank_proto_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bank_proto_bank_proto_goTypes = []interface{}{
	(*Card)(nil),                // 0: bank.Card
	(*ChargeItem)(nil),          // 1: bank.ChargeItem
	(*TransactionRequest)(nil),  // 2: bank.TransactionRequest
	(*ChargeId)(nil),            // 3: bank.ChargeId
	(*RefundRequest)(nil),       // 4: bank.RefundRequest
	(*TransactionResponse)(nil), // 5: bank.TransactionResponse
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_bank_proto_bank_proto_depIdxs = []int32{
	0, // 0: bank.ChargeItem.card:type_name -> bank.Card
	0, // 1: bank.TransactionRequest.card:type_name -> bank.Card
	2, // 2: bank.BankService.Charge:input_type -> bank.TransactionRequest
	4, // 3: bank.BankService.Refund:input_type -> bank.RefundRequest
	3, // 4: bank.BankService.GetCharge:input_type -> bank.ChargeId
	5, // 5: bank.BankService.Charge:output_type -> bank.TransactionResponse
	6, // 6: bank.BankService.Refund:output_type -> google.protobuf.Empty
	1, // 7: bank.BankService.GetCharge:output_type -> bank.ChargeItem
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bank_proto_bank_proto_init() }
func file_bank_proto_bank_proto_init() {
	if File_bank_proto_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bank_proto_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bank_proto_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bank_proto_bank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bank_proto_bank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bank_proto_bank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bank_proto_bank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bank_proto_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bank_proto_bank_proto_goTypes,
		DependencyIndexes: file_bank_proto_bank_proto_depIdxs,
		MessageInfos:      file_bank_proto_bank_proto_msgTypes,
	}.Build()
	File_bank_proto_bank_proto = out.File
	file_bank_proto_bank_proto_rawDesc = nil
	file_bank_proto_bank_proto_goTypes = nil
	file_bank_proto_bank_proto_depIdxs = nil
}
