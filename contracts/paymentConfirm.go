package contracts

type PaymentConfirmRequest struct {
}
type PaymentConfirmResponse struct {
	Status             int `json:"Status"`
	ConfirmationStatus int `json:"ConfirmationStatus"`
}
