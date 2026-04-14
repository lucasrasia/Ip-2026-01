package main
import "fmt"
func busca(banco[]int, v int)int{
	for _, i:=range banco{
		if banco[i]==v{
			return 1
		}
	}
	return -1
}
func main(){
	n := 9
	banco := make([]int, n)
	for i := 0; i < n; i++ {
    banco[i] = i + 1 // gera: [1, 2, 3, 4, 5, 6, 7, 8, 9]
}	
	var v int
	fmt.Print("Digite o valor: ")
	fmt.Scan(&v)
	b:=busca(banco, v)
	if b!=-1{
		fmt.Printf("%v está no banco de dados\n", v)
	} else{
		fmt.Printf("%v não está no banco de dados\n", v)
	}
}