package types

// ChargeRequest request sent in charge post api.
type ChargeRequest struct {
	Customer string   `json:"customer"`
	Card     CardItem `json:"card"`
	Merchant string   `json:"merchant"`
	Amount   float32  `json:"amount"`
}

// CardItem field from ChargeRequest.
type CardItem struct {
	Cvv            string `json:"cvv"`
	Number         string `json:"number"`
	ExpirationDate string `json:"expiration_date"`
	CardHolderName string `json:"card_holder_name"`
}
