package contracts

type PaymentResponse struct {
	MerchantOrderId string `json:"MerchantOrderId"`
	Payment         Payment
}

type Payment struct {
	Type               string `json:"Type"`
	SoftDescriptor     string `json:"SoftDescriptor"`
	PaymentDateTime    string `json:"PaymentDateTime"`
	Amount             int    `json:"Amount"`
	Installmentes      int    `json:"Installments"`
	Interest           string `json:"Interest"`
	Capture            bool   `json:"Capture"`
	ProductId          int    `json:"ProductId"`
	CreditCard         CreditCard
	PinPadInformation  PinPad
	PaymentId          string `json:"PaymentId"`
	Status             int    `json:"Status"`
	ConfirmationStatus int    `json:"ConfirmationStatus"`
}

type CreditCard struct {
	CardNumber                     string `json:"CardNumber"`
	ExpirationDate                 string `json:"ExpirationDate"`
	SecurityCodeStatus             string `json:"SecurityCodeStatus"`
	SecurityCode                   string `json:"SecurityCode"`
	BrandId                        int    `json:"BrandId"`
	IssuerId                       int    `json:"IssuerId"`
	InputMode                      string `json:"InputMode"`
	AuthenticationMethod           string `json:"AuthenticationMethod"`
	TruncateCardNumberWhenPrinting bool   `json:"TruncateCardNumberWhenPrinting"`
}

type PinPad struct {
	PhysicalCharacteristics string `json:"PhysicalCharacteristics"`
	ReturnDataInfo          string `json:"ReturnDataInfo"`
	SerialNumber            string `json:"SerialNumber"`
	TerminalId              string `json:"TerminalId"`
}
