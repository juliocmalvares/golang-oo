package contas

import (
	"fmt"
	"juliocmalvares/banco/clientes"
)

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int64
	saldo                    float64
}

func (c *ContaPoupanca) Saque(value float64) string {
	if c.saldo-value >= 0 && value >= 0 {
		c.saldo -= value
		return "Saque realizado"
	}
	return "Erro ao sacar"
}

func (c *ContaPoupanca) Deposito(value float64) (string, float64) {
	if value >= 0 {
		c.saldo += value
		return "Depósito realizado com sucesso, saldo: R$", c.saldo
	}
	return "Erro ao realizar depósito, saldo: R$", c.saldo
}

func (c *ContaPoupanca) Saldo() float64 {
	return c.saldo
}

func (co ContaPoupanca) Format(f fmt.State, c rune) {
	var value string = "--- Conta Poupança ---\n"
	value += "Titular: " + co.Titular.Nome + "\n"
	value += "Agencia: " + fmt.Sprintf("%d", co.Agencia) + "\n"
	value += "Conta: " + fmt.Sprintf("%d", co.Conta) + "\n"
	value += "saldo: R$ " + fmt.Sprintf("%.2f", co.saldo) + "\n"
	f.Write([]byte(value))
}
