package main
import "fmt"
func main(){
	var n int
	var r float64
	fmt.Print("Digite até onde vai a sequência: ")
	fmt.Scan(&n)
	if n<1 || n>333{
		fmt.Println("Digite um valor válido")
	} else{
		for i:=1; i<=n; i++{
			cima:=1003
			cima=cima-3
			if i%2!=0{
				r=r+float64(cima)/float64(i)
			} else{
				r=r-float64(cima)/float64(i)
			}
		}
		fmt.Println(r)
	}
}