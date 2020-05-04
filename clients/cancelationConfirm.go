package clients

import (
	f "fmt"
	contracts "physicalsales/contracts"

	resty "github.com/go-resty/resty/v2"
)

func CancelationConfirm(accessToken string, paymentId string) contracts.CancelationConfirmResponse {
	urlBase := "https://apisandbox.cieloecommerce.cielo.com.br"
	url := urlBase + "/1/physicalSales/" + paymentId + "/voids"

	payload := `{	
		"MerchantVoidId": "1587997297176",
		"MerchantVoidDate": "2020-04-27T14:21:37.176Z",
		"Card": {
		  "InputMode": "Typed",
		  "CardNumber": "5432123454321234"
		}
	  }`

	response := contracts.CancelationConfirmResponse{}

	client := resty.New()
	_, error := client.R().
		SetAuthToken(accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(&response).
		Post(url)

	if error != nil {
		f.Println("Houve alguma falha ao cancelar o pagamento", error)
	}
	return response
}
