package main
import "fmt"
func main(){
    var n int
    fmt.Print("Quantos alunos? ")
    fmt.Scan(&n)
    var soma float64
    status:="Aprovado"
    for i:=1; i<=n; i++{
        var n1, n2 float64
        fmt.Print("Digite as duas notas: ")
        fmt.Scan(&n1, &n2)
        media:=(n1+n2)/2
        if media<=3{
            status="Reprovado"
            fmt.Println("Média:", media, status)
        }else if media<=7{
            status="Exame"
            fmt.Println("Média:", media, status)
        }else{
            fmt.Println("Média:", media, status)
        }
        soma=soma+n1+n2
    }
    fmt.Println("Soma:", soma)
}

