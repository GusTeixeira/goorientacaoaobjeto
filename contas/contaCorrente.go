package contas

import (
	"github.com/GusTeixeira/goorientacaoaobjeto/clientes"
)

type ContaCorrente struct {
	Titular     clientes.Titular
	NumeroConta int
	saldo       float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return ("Saque realizado com sucesso")
	} else {
		return ("saldo Inválido ou Insuficiente")
	}

}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {

	if valorDoDeposito >= 0 {
		c.saldo += valorDoDeposito
		return ("Depósito efetuado")
	}

	return ("Depósito não efetuado")

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	if valorDaTransferencia <= 0 || valorDaTransferencia >= c.saldo {
		return ("Valor Inválido ou Insuficiente")
	}

	c.saldo -= valorDaTransferencia
	contaDestino.Depositar(valorDaTransferencia)
	return ("Valor transferido com sucesso")
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
