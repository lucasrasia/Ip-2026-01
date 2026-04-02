package main
import "fmt"

func main(){
	var v float64
	var total float64
	fmt.Print("Valor inicial do carro: ")
	fmt.Scan(&v)
	total+=v
	var n1 string
	fmt.Print("Deseja Ar condicionado(s/n): ")
	fmt.Scan(&n1)
	if n1=="s"{
		total+=1750
	}
	var n2 string
	fmt.Print("Deseja Pintura metálica(s/n): ")
	fmt.Scan(&n2)
	if n2=="s"{
		total+=800
	}
	var n3 string
	fmt.Print("Vidro elétrico(s/n): ")
	fmt.Scan(&n3)
	if n3=="s"{
		total+=1200
	}
	var n4 string
	fmt.Print("Direção hidráulica(s/n): ")
	fmt.Scan(&n4)
	if n4=="s"{
		total+=2000
	}		
	fmt.Printf("Valor total do carro: %v\n", total)
}
