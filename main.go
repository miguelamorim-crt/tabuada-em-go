package main

import "fmt"

func tabuada() {
	fmt.Println("tabuada em Go!")

	fmt.Println("a tabuada e do numero:")
	var numero int
	fmt.Scan(&numero)

	fmt.Println("quantas vezes quer repetir:")
	var quantidade int
	fmt.Scan(&quantidade)

	for i := 1; i <= quantidade; i++ {
		fmt.Printf("%d X %d = %d\n", numero, i, i*numero)
	}
}

func menu() {
	fmt.Println("=== tabuada Go! ===")
	fmt.Println("1 - tabuada")
	fmt.Println("0 - sair")
}

func main() {
	var opcao int
	for {
		menu()
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			tabuada()
		case 0:
			return
		default:
			fmt.Println("opcao invalida!")
		}
	}
}
