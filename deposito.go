package banco


type Conta struct {
	numero string
	saldo  float64
}

func NovaConta(num string) *Conta {
	return &Conta{
		numero: num,
		saldo:  0.0,
	}
}

func (c *Conta) Despositar(aumento float64) {
	c.saldo += aumento
}

func (c Conta) Saldo() float64 {
	return c.saldo
}

func (c *Conta) RetirarDinheiro(retirar float64){
	c.saldo -= retirar
}
