package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Charge struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	From            string             `bson:"from"`
	Card            CardItem           `bson:"card"`
	To              string             `bson:"to"`
	Amount          float32            `bson:"amount"`
	TransactionType string             `bson:"transaction_type"`
}
