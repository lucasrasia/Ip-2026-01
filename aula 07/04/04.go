package main
import "fmt"
func fat(n int)int{
	 if n==0 || n==1{
		return 1
	 }
	 return n*fat(n-1)
}
func main(){
	var n int
	fmt.Print("Número: ")
	fmt.Scan(&n)
	f:=fat(n)
	fmt.Printf("%v fatorial é %v\n", n, f)
}
