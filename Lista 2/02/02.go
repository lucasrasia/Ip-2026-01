package main
import("fmt")

func main(){
	var n int
	fmt.Print("Número: ")
	fmt.Scan(&n)
	if n>0 {
		fmt.Printf("%v é positivo \n", n)		
	} else if n<0 {
		fmt.Printf("%v é negativo \n", n)
	} else {
		fmt.Printf("%v é nulo \n", n)
	}
}
