package contas

import (
	"fmt"
	c "juliocmalvares/banco/clientes"
)

type ContaCorrente struct {
	Titular       c.Titular
	NumeroAgencia int32
	NumeroConta   int32
	saldo         float64
}

func (c *ContaCorrente) Saque(value float64) string {
	if c.saldo-value >= 0 && value >= 0 {
		c.saldo -= value
		return "Saque realizado"
	}
	return "Erro ao sacar"
}

func (c *ContaCorrente) Deposito(value float64) (string, float64) {
	if value >= 0 {
		c.saldo += value
		return "Depósito realizado com sucesso, saldo: R$", c.saldo
	}
	return "Erro ao realizar depósito, saldo: R$", c.saldo
}

func (c *ContaCorrente) Transferir(value float64, conta *ContaCorrente) string {
	if value >= 0 && c.saldo-value >= 0 {
		c.saldo -= value
		conta.saldo += value
		return "Transferência realizada com sucesso"
	}
	return "Erro ao fazer transferência"
}

func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}

// funcão de format para print

func (co ContaCorrente) Format(f fmt.State, c rune) {
	var value string = ""
	value += "Titular: " + co.Titular.Nome + "\n"
	value += "Agencia: " + fmt.Sprintf("%d", co.NumeroAgencia) + "\n"
	value += "Conta: " + fmt.Sprintf("%d", co.NumeroConta) + "\n"
	value += "saldo: R$ " + fmt.Sprintf("%.2f", co.saldo) + "\n"
	f.Write([]byte(value))
}
