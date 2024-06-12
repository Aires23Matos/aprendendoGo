package main

import "fmt"

func main() {
	nome := []string{
		"Aires, ",
		"Bernado, ",
		"Pedro",
	}

	var p1 []string = nome[0: 2]
	p1 [0]="Modificar" 
	fmt.Print(p1)
}
