package main
import "fmt"
func main() {
	var alturas [10]float64
	soma := 0.0
	for i := 0; i < 10; i++ {
		fmt.Scan(&alturas[i])
		soma += alturas[i]
	}
	media := soma / 10
	for i := 0; i < 10; i++ {
		if alturas[i] > media {
			fmt.Println(alturas[i])
		}
	}
}
