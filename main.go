package main

import (
	f "fmt"
	"physicalsales/clients"
	"strconv"
)

func main() {

	//Gera o token
	token := clients.GetToken().AccessToken
	f.Println("O token gerado foi : ", token)

	//Gera o Pagamento
	generatePayment := clients.CreatePayment(token)
	f.Println("O Pagamento foi gerado e o paymentId é : " + generatePayment.Payment.PaymentId)

	//Confirma o Pagamento
	confirm := clients.PaymentConfirm(token, generatePayment.Payment.PaymentId)
	confirmStatus := strconv.Itoa(confirm.Status)
	f.Println("O Pagamento foi confirmado e o status é : " + confirmStatus)

	//Cancela o Pagamento
	cancelation := clients.CancelationConfirm(token, generatePayment.Payment.PaymentId)
	cancelationStatus := strconv.Itoa(cancelation.Status)
	f.Println("O Cancelamento do pagamento foi confirmado e o status é : " + cancelationStatus)

	//Desfaz o Cancelamento
	undoCancelation := clients.UndoCancelation(token, generatePayment.Payment.PaymentId, cancelation.VoidId)
	undoCancelationStatus := strconv.Itoa(undoCancelation.Status)
	f.Println("O Status do Desfazimento do Cancelamento de Pagamento é: " + undoCancelationStatus)

}
