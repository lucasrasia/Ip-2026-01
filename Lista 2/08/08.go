package main
import ("fmt")

func main(){
	var n float64
	fmt.Print("Digíte um número: ")
	fmt.Scan(&n)
	if 20<n && n<90 {
		fmt.Printf("%v está entre 20 e 90 \n", n)
	} else{
		fmt.Printf("%v não está entre 20 e 90 \n", n)
	}
}
