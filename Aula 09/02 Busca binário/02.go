package main
import "fmt"

func busca(banco []int, v int) int {
	i := 0
	fim := len(banco) - 1
	for i <= fim {
		meio := (i + fim) / 2
		if banco[meio] == v {
			return 1
		} else if banco[meio] > v {
			fim = meio - 1
		} else {
			i = meio + 1
		}
	}
	return -1
}
func main() {
	banco := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var v int
	fmt.Print("Digíte o valor: ")
	fmt.Scan(&v)
	b := busca(banco, v)
	if b != -1 {
		fmt.Printf("%v está no banco de dados\n", v)
	} else {
		fmt.Printf("%v não está no banco de dados\n", v)
	}
}
