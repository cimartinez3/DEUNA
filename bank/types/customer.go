package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// CustomerItem struct that defines customer entity in mongo.
type CustomerItem struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	CustomerId string             `bson:"customer_id,omitempty"`
	Card       CardItem           `bson:"card"`
	Balance    float32            `bson:"balance"`
}
