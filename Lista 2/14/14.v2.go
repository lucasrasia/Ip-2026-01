package main
import "fmt"
func main(){
	opcional:=[]struct{
		nome string
		valor float64
	}{
		{"Ar condicioando", 1750},
		{"Pintura metálica", 800},

	}
	var n float64
	fmt.Print("Valor inicial do carro: ")
	fmt.Scan(&n)
	total:=n
	for _, op:=range opcional{
		var sn string
		fmt.Printf("Desseja %v? (s/n) ", op.nome)
		fmt.Scan(&sn)
		if sn=="s"{
			total=total + op.valor
		}
	}
	fmt.Printf("O valor final do carro é: R$%.2f\n", total)

}
