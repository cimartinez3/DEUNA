package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Merchant struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Balance float64            `bson:"balance"`
}
