package main
import "fmt"
func main(){
	a:=[]int{6,4,7,4,9,3,6,23}
	for i:=0; i<len(a)-1; i++{
		for j:=0; j<len(a)-1-i; j++{ //a cada passada o último termo é colocado na ordem certa
			if a[j]>a[j+1]{
				a[j], a[j+1]=a[j+1], a[j]  // troca simultaneamente
			}
		}
	}
	fmt.Println(a)
}
