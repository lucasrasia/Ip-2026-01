package main
import "fmt"
type pessoa struct{
	NOME string
	ALTURA float64
	PESO float64
}

func main(){
	var ind []pessoa //ver isso
	var nome string
	var h float64
	for {
		fmt.Print("Nome (digíte 'FIM' para sair): ")
		fmt.Scanln(&nome)
		if nome =="FIM" {
			break
		} 
		fmt.Print("Altura: ")
		fmt.Scanln(&h)
		peso_ideal:=72.7*h-58
		fmt.Printf("Peso ideal: %.2f\n\n", peso_ideal)
		p := pessoa{
			NOME : nome,
			ALTURA : h, 
			PESO : peso_ideal,
		}
		ind=append(ind, p) //ver isso
	}
	fmt.Println("\nIndivíduos cadastrados:")
	for i := range ind {  //ver isso
		fmt.Printf("Nome: %v, Altura: %v, Peso ideal: %.2f\n", ind[i].NOME, ind[i].ALTURA, ind[i].PESO)
	}
}
