package banco_test

import (
	banco "estando"
	"testing"
)

func TestDeposito(t *testing.T) {
	//Arrange
	conta := banco.NovaConta("12345")
	iniciarconta := conta.Saldo()
	saldodepositar := 100.00
	//Act
	conta.Despositar(saldodepositar)
	//Assert
	saldoesperado := saldodepositar + iniciarconta
	saldoatual := conta.Saldo()
	if saldoatual != saldoesperado {
		t.Errorf("não foi feito o depósito corretamente: seu saldo %v , saldo a ser depositado %v", saldoatual, saldodepositar)
	}
}

func TestRetirarDinheiro(t *testing.T) {
	//Arrenge
	conta := banco.NovaConta("12345")
	conta.Despositar(200.00)
	iniciarconta := conta.Saldo()
	saldoretirar := 100.00
	//Act
	conta.RetirarDinheiro(saldoretirar)
	//Assert
	saldoretirado := iniciarconta - saldoretirar
	saldoatual := conta.Saldo()
	if saldoretirado > 0{
		if saldoretirado != saldoatual {
			t.Errorf("Erro esperado do saldo atual %v rebendo do saldo atual %v", saldoatual, saldoretirado)
		}
	}else {
		if saldoretirado >= saldoatual {
			t.Errorf("Saldo insuficiente: saldo da conta %v pedido para retirar %v", iniciarconta, saldoretirado)
		}
	}
}
