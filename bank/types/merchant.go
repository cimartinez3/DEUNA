package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// Merchant struct that defines Merchant entity in mongo.
type Merchant struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Balance float64            `bson:"balance"`
}
