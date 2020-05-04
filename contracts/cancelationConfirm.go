package contracts

type CancelationConfirmResponse struct {
	VoidId             string `json:"VoidId"`
	CancellationStatus int    `json:"CancellationStatus"`
	ConfirmationStatus int    `json:"ConfirmationStatus"`
	Status             int    `json:"Status"`
}
