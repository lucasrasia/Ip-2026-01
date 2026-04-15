package main
import "fmt"
func fat(v int)int{
	if v==0 || v==1{
		return 1
	}
	return v*fat(v-1)
}
func main(){
	var x float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&x)
	var soma float64
	for i:=2; i<=20; i++{
		f:=fat(i)
		if i%2==0{
			soma=soma+(x/float64(f))
		}else{
			soma=soma-(x/float64(f))
		}
	}
	fmt.Printf("Resultado: %.2f\n", soma)
}