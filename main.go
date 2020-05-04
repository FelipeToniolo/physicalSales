package main

import (
	f "fmt"
	"physicalsales/clients"
)

func main() {

	token := clients.GetToken().AccessToken

	if token != "" {
		f.Println("O token gerado foi : ", token)
	}

}
