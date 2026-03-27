package main
import("fmt")

var n1 int
var n2 int
func main(){
	fmt.Print("Número 1: ")
	fmt.Scan(&n1)
	fmt.Print("Número 2: ")
	fmt.Scan(&n2)
	n:=n1+n2
	if n>20 {
		fmt.Println(n+8)
	} else {
		fmt.Println(n-5)
	}
}