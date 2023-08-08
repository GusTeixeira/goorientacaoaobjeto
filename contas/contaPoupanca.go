package contas

import (
	"github.com/GusTeixeira/goorientacaoaobjeto/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {

	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return ("Saque realizado com sucesso")
	} else {
		return ("saldo Inválido ou Insuficiente")
	}

}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) string {

	if valorDoDeposito >= 0 {
		c.saldo += valorDoDeposito
		return ("Depósito efetuado")
	}

	return ("Depósito não efetuado")

}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) string {
	if valorDaTransferencia <= 0 || valorDaTransferencia >= c.saldo {
		return ("Valor Inválido ou Insuficiente")
	}

	c.saldo -= valorDaTransferencia
	contaDestino.Depositar(valorDaTransferencia)
	return ("Valor transferido com sucesso")
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
