package main
import "fmt"
func main(){
	for i:=1; i<=10; i++{
		for j:=i; j<=i; j++{
			fmt.Printf("[%v, %v] ", i, j)
		}
		fmt.Println()
	}
}