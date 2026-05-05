package main
import "fmt"
func main(){
	var F[]int
	var f int
	F=append(F, 0)
	F=append(F, 1)
	for i:=0;i<=47;i++{
		f=F[i]+F[i+1]
		F=append(F, f)
	}
	fmt.Println(F)
}
