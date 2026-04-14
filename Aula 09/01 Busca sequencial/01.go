package main
import "fmt"
func main(){
	n := 9
array := make([]int, n)

for i := 0; i < n; i++ {
    array[i] = i + 1 // gera: [1, 2, 3, 4, 5, 6, 7, 8, 9]
}
	var n int
	fmt.Print("Digite o número para buscar: ")
	fmt.Scan(&n)
	frase:=false
	for _, i:=range banco{
		if i==n{
			frase=true
		} 
	}
	if frase{
		fmt.Printf("%v está no banco\n", n)
	} else{
		fmt.Printf("%v não está no banco\n", n)
	}
}
