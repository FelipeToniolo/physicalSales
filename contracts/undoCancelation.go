package contracts

type UndoCancelationResponse struct {
	CancellationStatus int `json:"CancellationStatus"`
	ConfirmationStatus int `json:"ConfirmationStatus"`
	Status             int `json:"Status"`
}
