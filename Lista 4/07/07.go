package main
import "fmt"
func main(){
	var A[]int
	for i:=1;i<=99;i+=2{
		A=append(A, i)
	}
	fmt.Println(A)
}

