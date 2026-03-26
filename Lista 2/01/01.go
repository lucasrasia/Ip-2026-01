package main
import("fmt")

func main(){
	var n int
	fmt.Print("Digíte o número: ")
	fmt.Scan(&n)
	if n%2==0 {
		fmt.Printf("%v é par\n", n)
	}else {
		fmt.Printf("%v é ímpar\n", n)
	}
}