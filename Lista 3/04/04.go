package main
import "fmt"
func main(){
	var n int
	for{
		fmt.Print("Digíte um quadrado perfeito: ")
		fmt.Scan(&n)
		if n<=0{
			return 
		}
		var qp=false
		for i:=1; i*i<=n; i++{
			if i*i==n{
				qp=true
			}
		}
		if qp{
			fmt.Printf("%v é quadrado perfeito! \n\n", n)
		} else{
			fmt.Printf("%v NÃO é quadrado perfeito! \n\n", n)
		}
}
}