package main
import "fmt"

func main(){
	mes:=[12]string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}
	var d, m, a int
	fmt.Print("Coloque uma data(dd mm aaaa): ")
	fmt.Scan(&d, &m, &a)
	fmt.Printf("%v de %v de %v \n", d, mes[m-1], a)
}