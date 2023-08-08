package main

import (
	"fmt"

	"github.com/GusTeixeira/goorientacaoaobjeto/clientes"
	"github.com/GusTeixeira/goorientacaoaobjeto/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "121220120",
		Profissao: "Programador",
	}
	contaDoBruno := contas.ContaCorrente{
		Titular:     clienteBruno,
		NumeroConta: 123,
	}

	clienteDenis := clientes.Titular{
		Nome:      "Denis",
		CPF:       "121220120",
		Profissao: "Gay",
	}
	contaDoDenis := contas.ContaCorrente{
		Titular:     clienteDenis,
		NumeroConta: 123,
	}

	fmt.Println(contaDoBruno)
	contaDoBruno.Depositar(100)
	fmt.Println(contaDoBruno.ObterSaldo())
	PagarBoleto(&contaDoBruno, 20)
	fmt.Println(contaDoBruno.ObterSaldo())

	fmt.Println(contaDoDenis)
	contaDoDenis.Depositar(100)
	fmt.Println(contaDoDenis.ObterSaldo())
	PagarBoleto(&contaDoDenis, 20)
	fmt.Println(contaDoDenis.ObterSaldo())

}
