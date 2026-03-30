package main
import ("fmt")

var r int
var i int
func main(){
	fmt.Print("Norte(1), Nordeste(2), Centro-Oeste(3), Sul(4): ")
	fmt.Scan(&r)
	fmt.Print("Ida(1) ou Ida e volta(2): ")
	fmt.Scan(&i)
	if r>=5 && r<0 && i!=1 && i!=2 {
		fmt.Println("Região ou ida/volta inválida")
	} else if r==1 {
		a:=500
		if i==2{
			b:=a+400
			fmt.Printf("Valor: %v \n", b)
		} else {
			fmt.Printf("Valor %v \n", a)
		}
	} else if r==2 {
		a:=350
		if i==2{
			b:=a+300
			fmt.Printf("Valor: %v \n", b)
		} else {
			fmt.Printf("Valor %v \n", a)
		}
	} else if r==3 {
		a:=350
		if i==2{
			b:=a+250
			fmt.Printf("Valor: %v \n", b)
		} else {
			fmt.Printf("Valor %v \n", a)
		}
	} else{
		a:=300
		if i==2{
			b:=a+250
			fmt.Printf("Valor: %v \n", b)
		} else {
			fmt.Printf("Valor %v \n", a)
		}
	}
	
}