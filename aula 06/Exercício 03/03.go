package main
import "fmt"
func main() {
    var nota[10]float64
    for i:=0; i<10; i++ {
        fmt.Printf("Nota %v: ", (i+1))
        fmt.Scan(&nota[i])
    }
    for i := range nota {
        fmt.Printf("%v ", nota[(9-i)])
    }
}
