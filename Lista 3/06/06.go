package main
import "fmt"

func main() {
  var n int
  fmt.Print("Digite um número: ")
  fmt.Scan(&n)
  e_triangular:=false
  for i:=3; i<=n; i++ {
      if i*(i-1)*(i-2)==n{
          e_triangular=true
      }
  }
  if e_triangular{
      fmt.Printf("%v é triangular", n)
  } else{
      fmt.Print("%v não é triangular", n)
  }
}
