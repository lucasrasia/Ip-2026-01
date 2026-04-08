package main
import ("fmt")
func verificar(n int) {
    if n < 0 {
        fmt.Println("Número negativo!")
        return  // sai da função sem retornar valor
    }
    fmt.Println("Número válido")
}
func main(){
	var n int
	fmt.Print(" ")
	fmt.Scan(&n)
	f:=verificar(n)
	fmt.Println(f)
}