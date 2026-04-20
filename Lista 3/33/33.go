package main
import "fmt"
func main(){
	var n, d int
	fmt.Print("Numerador: ")
	fmt.Scan(&n)
	fmt.Print("Denominador: ")
	fmt.Scan(&d)
	if d>n{
		fmt.Printf("Quociente(%v,%v): 0, Resto(%v,%v): %v\n", n, d, n, d, n)
	} else if d==0{
		fmt.Println("Digite um denominador válido")
	} else if d==n{
		fmt.Printf("Quociente(%v,%v): 1, Resto(%v,%v): 0\n", n, d, n, d)
	} else{
		r:=n
		var q int
		for{
			if r<d{
				break
			}
			r=r-d
			q++
		}
		fmt.Printf("Quociente(%v,%v): %v, Resto(%v,%v): %v\n", n, d,q,n,d,r)
	}
}