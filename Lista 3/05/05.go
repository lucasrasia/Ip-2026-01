package main
import "fmt"
type pessoa struct{
	idade int
	altura float64
	peso float64
}
func main(){
	var ind[]pessoa
	var i int
	var n string
	var h, p float64
	for{
		fmt.Print("Deseja continuar? ('s' ou 'n') ")
		fmt.Scan(&n)
		if n=="n"{
			return
		}
		fmt.Print("Idade: ")
		fmt.Scan(&i)
		fmt.Print("Altura: ")
		fmt.Scan(&h)
		fmt.Print("Peso: ")
		fmt.Scan(&p)
		p:=pessoa{
			idade: i,
			altura: h,
			peso: p,
		}
		ind=append(ind, p)
	}
	var a int
	for i:=range ind {
		if ind[i].idade>50{
			a=a+1
		}
	}
}
