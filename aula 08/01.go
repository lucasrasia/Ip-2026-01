package main
import "fmt"
type pessoa struct{
	nome string
	altura float64
}

func main(){
	var ind []pessoa
	var name string
	var h float64
	for {
		fmt.Println("Para parar digite 'FIM' ")
		fmt.Print("Nome: ")
		fmt.Scan(&name)
		if name =="FIM" {
			break
		} 
		fmt.Print("Altura: ")
		fmt.Scan(&h)
		p := pessoa{
			nome : name,
			altura : h, 
		}
		ind=append(ind, p)
	}
}