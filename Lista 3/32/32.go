package main
import "fmt"
func main(){
	var a, b int
	fmt.Print("N1: ")
	fmt.Scan(&a)
	fmt.Print("N2: ")
	fmt.Scan(&b)
	var r int
	for i:=1; i<=b; i++{
		r+=a
	}
	fmt.Println(r)
}