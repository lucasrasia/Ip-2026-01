package main
import "fmt"
func fat(v int)int{
	if v==0 || v==1{
		return 1
	}
	return v*fat(v-1)
}
func main(){
	var v int
	fmt.Print("Digite um número: ")
	_, err:=fmt.Scan(&v)
	if err!=nil || v<0{
		fmt.Println("Digite um valor válido")

	}
	f:=fat(v)
	fmt.Printf("%v fatorial %v\n", v, f)
}