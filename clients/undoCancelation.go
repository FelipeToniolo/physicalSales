package clients

import (
	f "fmt"
	contracts "physicalsales/contracts"

	resty "github.com/go-resty/resty/v2"
)

func UndoCancelation(accessToken string, paymentId string, voidId string) contracts.UndoCancelationResponse {
	urlBase := "https://apisandbox.cieloecommerce.cielo.com.br"
	url := urlBase + "/1/physicalSales/" + paymentId + "/voids" + voidId

	body := "{}"

	response := contracts.UndoCancelationResponse{}

	client := resty.New()
	_, error := client.R().
		SetAuthToken(accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Delete(url)

	if error != nil {
		f.Println("Houve alguma falha ao desfazer o cancelamento do pagamento", error)
	}

	return response
}
