package main
import "fmt"
func main() {
    var n[5]int
    var soma int
    for i:=0; i<5; i++ {
        fmt.Printf("Número %v: ", (i+1))
        fmt.Scan(&n[i])
        soma+=n[i]
    }
    fmt.Println(soma)
}
