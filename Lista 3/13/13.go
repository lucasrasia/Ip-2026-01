package main
import "fmt"

func main() {
	a := 1.0
	b := 1.0
	h := 1.0
	for i := 1; i < 50; i++ {
		h = h + (b/a)
		a++
		b+=2
	}
	fmt.Printf("O valor de h é: %.3f\n", h)
}
