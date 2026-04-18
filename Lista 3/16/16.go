package main
import "fmt"
func main(){
	var a[]int
	var x, y, z int
	var n int
	fmt.Print("Quantidade de termos: ")
	fmt.Scan(&n)
	fmt.Print("Primeiros termos: ")
	fmt.Scan(&x, &y)
	a=append(a, x, y)
	for i:=2; i<n; i++{
		if i%2==0{
			z=a[i-1]-a[i-2]
			a=append(a, z)
		} else{
			z=a[i-1]+a[i-2]
			a=append(a, z)
		}
	}
	fmt.Println("Série de Fetuccine:")
	for _, termo := range a {
		fmt.Printf("%v ", termo)
	}
	fmt.Println()
}