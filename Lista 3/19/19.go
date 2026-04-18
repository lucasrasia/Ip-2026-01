package main
import "fmt"
func main(){
	for i:=1; i<=10; i++{
		for j:=i+1; j<=10; j++{
			fmt.Printf("[%v, %v] ", i, j)
		}
		fmt.Println()
	}
}
