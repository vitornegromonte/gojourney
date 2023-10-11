package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A implementação de um laço for requer que seja seguida uma certa estrutura:
	// inicialização:  j := 7;
	// condição: j <= 9;
	// execução pós iteração: j++ (soma +1 em j para cada iteração de laço)
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
