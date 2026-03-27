package main
import ("fmt"; "math")

func main(){
	var n float64
	fmt.Print("Númrero: ")
	fmt.Scan(&n)
	if n>=0 {
		conta:= math.Sqrt(n)
		fmt.Printf("%.2f \n", conta)
	} else {
		conta:= math.Pow(n, 2)
		fmt.Println(conta)
	}
}
