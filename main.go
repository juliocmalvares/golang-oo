package main

import (
	"fmt"
	cl "juliocmalvares/banco/clientes"
	co "juliocmalvares/banco/contas"
)

func main() {
	contaDoJulio := co.ContaCorrente{
		Titular: cl.Titular{
			Nome:      "Julio",
			Cpf:       "132.123.132-12",
			Profissao: "Desenvolvedor"},
		NumeroAgencia: 6882,
		NumeroConta:   504019,
	}

	fmt.Println(contaDoJulio)

	contaP := co.ContaPoupanca{}
	fmt.Println(contaP)
}
