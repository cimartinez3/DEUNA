package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// CardItem struct that defines card entity in mongo.
type CardItem struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	Cvv            string             `bson:"cvv"`
	Number         string             `bson:"number"`
	ExpirationDate string             `bson:"expiration_date"`
	CardHolderName string             `bson:"card_holder_name"`
}
