package contracts

type PaymentConfirmResponse struct {
	Status             int `json:"Status"`
	ConfirmationStatus int `json:"ConfirmationStatus"`
}
