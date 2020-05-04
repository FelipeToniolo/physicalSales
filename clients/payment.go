package clients

import (
	f "fmt"
	contracts "physicalsales/contracts"

	resty "github.com/go-resty/resty/v2"
)

func CreatePayment(accessToken string) contracts.PaymentResponse {

	urlBase := "https://apisandbox.cieloecommerce.cielo.com.br"
	url := urlBase + "/1/physicalSales"

	response := contracts.PaymentResponse{}

	payload := `{
		"MerchantOrderId": "1587997030607",
			"Payment": {
			"Type": "PhysicalCreditCard",
				"SoftDescriptor": "Desafio GO 2",
				"PaymentDateTime": "2020-01-08T11:00:00",
				"Amount": 100,
				"Installments": 1,
				"Interest": "ByMerchant",
				"Capture": true,
				"ProductId": 1,
				"CreditCard": {
				"CardNumber": "5432123454321234",
					"ExpirationDate": "12/2021",
					"SecurityCodeStatus": "Collected",
					"SecurityCode": "123",
					"BrandId": 1,
					"IssuerId": 401,
					"InputMode": "Typed",
					"AuthenticationMethod": "NoPassword",
					"TruncateCardNumberWhenPrinting": true
			},
			"PinPadInformation": {
				"PhysicalCharacteristics": "PinPadWithChipReaderWithoutSamAndContactless",
					"ReturnDataInfo": "00",
					"SerialNumber": "0820471929",
					"TerminalId": "42004558"
			}
		}
	}`

	client := resty.New()
	_, error := client.R().
		SetAuthToken(accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(&response).
		Post(url)

	if error != nil {
		f.Println("Houve alguma falha ao criar o pagamento", error)
	}

	return response

}
