package clients

import (
	f "fmt"
	contracts "physicalsales/contracts"

	resty "github.com/go-resty/resty/v2"
)

func PaymentConfirm(accessToken string, paymentId string) contracts.PaymentConfirmResponse {
	urlBase := "https://apisandbox.cieloecommerce.cielo.com.br"
	url := urlBase + "/1/physicalSales/" + paymentId + "/confirmation"

	body := "{}"

	response := contracts.PaymentConfirmResponse{}

	client := resty.New()
	_, error := client.R().
		SetAuthToken(accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Put(url)

	if error != nil {
		f.Println("Houve alguma falha ao confirmar o pagamento", error)
	}

	return response
}
