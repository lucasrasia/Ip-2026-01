package main
import "fmt"

func main(){
	a:=[]int{6,4,7,4,9,3,6,23}
	for i:=1; i<len(a); i++{
		for j:=0; j<i; j++{
			if a[j]>a[i]{
				a[j], a[i]=a[i], a[j]
			}
		}
	}6,4,7,4,9,3,6,23
	fmt.Println(a)
}