package main

import (
	"fmt"
	cl "juliocmalvares/banco/clientes"
	co "juliocmalvares/banco/contas"
)

func PagarBoleto(conta verificaConta, value float64) {
	conta.Saque(value)
}

type verificaConta interface {
	Saque(value float64) string
}

func main() {
	contaDoJulio := co.ContaCorrente{
		Titular: cl.Titular{
			Nome:      "Julio",
			Cpf:       "132.123.132-12",
			Profissao: "Desenvolvedor"},
		NumeroAgencia: 6882,
		NumeroConta:   504019,
	}
	fmt.Println("Antes do boleto")
	contaDoJulio.Deposito(1000)
	fmt.Println(contaDoJulio)

	contaP := co.ContaPoupanca{}
	contaP.Deposito(1000)
	fmt.Println(contaP)

	fmt.Println("Depois do boleto")
	PagarBoleto(&contaDoJulio, 129.99)
	PagarBoleto(&contaP, 129.99)

	fmt.Println(contaDoJulio)
	fmt.Println(contaP)
}
