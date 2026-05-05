package main
import "fmt"
func main(){
	a:=[4]int{7, 6, 34, 11}
	for i:=0; i<len(a)-1; i++{
		for j:=0; j<len(a)-1-i; i++{
			if a[i]>a[i+1]{
				a[j], a[j+1]=a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}
