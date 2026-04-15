package main
import "fmt"

func main(){
	var n int
	fmt.Print("Digite até onde deseja ir: ")
	fmt.Scan(&n)
	
	n1 := 0
	n2 := 1
	
	fmt.Printf("%v ", n1)
	fmt.Printf("%v ", n2)
	
	for i := 3; i <= n; i++ {
		n3 := n1 + n2
		fmt.Printf("%v ", n3)
		n1 = n2
		n2 = n3
	}
	fmt.Println()
}