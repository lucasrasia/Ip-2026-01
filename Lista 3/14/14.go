package main
import "fmt"
func main(){
	var n1, n2 int
	fmt.Print("Digite os valores n1 e n2: ")
	fmt.Scan(&n1, &n2)
	max:=max(n1, n2)        
	min:=min(n1, n2)
	for i:=min; i<=max; i++{    
		primo:=true
		for n:=2; n<i; n++{
			if i%n==0 {
				primo=false
				break
			}
		}
		if primo{
			fmt.Println(i)
		}
	}
}