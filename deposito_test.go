package banco_test

import (
	"testing"
	banco "estando"
)

func TestDeposito(t* testing.T){
	//Arrange
	conta := banco.NovaConta("12345")
	iniciarconta := conta.Saldo()
	saldodepositar := 100.00
	//Act
	conta.Despositar(float32(saldodepositar))
	//Assert
	saldoesperado := saldodepositar + iniciarconta
	saldoatual := conta.Saldo()
	if saldoatual != saldoesperado {
		t.Errorf("não foi feito o depósito corretamente: seu saldo %v , saldo a ser depositado %v",saldoatual,saldodepositar )
	}
}