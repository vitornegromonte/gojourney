// O go não conta com a estrutura de repetição while, porém,
// a mesma pode ser implementada utilizando um laço for
// é importante ressaltar que ela não é exatamente uma implementação nativa
// ela apenas se comporta de um laço while em outras linguagens

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// Função apenas para desmontrar um laço infinito
func infinity() {
	for {

	}
}
