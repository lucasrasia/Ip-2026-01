package main
import("fmt")

func main(){
	var c float64
	fmt.Print("Insira o valor da compra: ")
	fmt.Scan(&c)
	if c<=0 {
		fmt.Println("Valor de compra inválido")
	} else if c<10 {
		l:=c*0.7
		fmt.Printf("Valor da venda: %.2f \n", l)
	} else if c<30 {
		l:=c*0.5
		fmt.Println("Valor da venda:", l)
	} else if c<50 {
		l:=c*0.4
		fmt.Println("Valor da venda:", l)
	} else {
		l:=c*0.3
		fmt.Println("Valor da venda:", l)
	}
}
