package clients

import (
	f "fmt"
	contracts "physicalsales/contracts"
	credential "physicalsales/credentials"

	resty "github.com/go-resty/resty/v2"
)

func GetToken() contracts.AuthResponse {

	response := contracts.AuthResponse{}

	urlBase := "https://authsandbox.braspag.com.br"
	user := credential.GetUser()
	password := credential.GetPassword()
	url := urlBase + "/oauth2/token"

	client := resty.New()
	_, error := client.R().
		SetBasicAuth(user, password).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody("grant_type=client_credentials").
		SetResult(&response).
		Post(url)

	if error != nil {
		f.Println("Falha na Geraçao do Token")
	}
	return response
}
