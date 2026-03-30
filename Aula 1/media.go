package main
import ("fmt")

func main(){
	var n int
	var nota float32
	var soma float32
	fmt.Print("Número de alunos: ")
	fmt.Scan(&n)
	for i:=1; i < (n+1); i++ {
		fmt.Printf("Nota %v: ", i)
		fmt.Scan(&nota)
		soma+=nota
	}
	media:=(soma)/float32(n)
	fmt.Printf("Média: %.2f \n", media)
}
