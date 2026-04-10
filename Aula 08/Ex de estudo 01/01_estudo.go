package main
import "fmt"
type carro struct {
    marca string
    ano int
    preco float64
}

func main() {
    var car[]carro
    var m string
    var a int
    var p float64
    for {
    fmt.Print("Digíte a marca (para parar digitem 'FIM'): ")
    fmt.Scan(&m)
    if m=="FIM"{
        break
    }
    fmt.Print("Digíte o ano: ")
    fmt.Scan(&a)
    fmt.Print("Digíte o preço: ")
    fmt.Scan(&p)
    c:=carro{
        marca: m,
        ano: a,
        preco: p,
    }
    car=append(car, c)
    }
    fmt.Println("\n====Carros desde 2020====")
    for i:=range car {
        if car[i].ano>2020{
            fmt.Printf("marca: %v, preço: %v, ano: %v \n", car[i].marca, car[i].preco, car[i].ano)
        }
    }
