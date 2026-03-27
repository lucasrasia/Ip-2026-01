package main
import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Número A: ")
	fmt.Scan(&a)
	fmt.Print("Número B: ")
	fmt.Scan(&b)
	if a%b == 0 {
		fmt.Printf("%v é divisível por %v \n", a, b)
	} else{
		fmt.Printf("%v não é divisível por %v \n", a, b)
	}
}

