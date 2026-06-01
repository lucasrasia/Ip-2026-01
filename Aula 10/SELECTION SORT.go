package main
import "fmt"

func main(){
	a:=[]int{3,4,6,8,5,3,67,9}
	for i:=0; i<len(a)-1; i++{
		menor:=a[i]
		for j:=i+1; j<len(a); j++{
			if j<menor{
				menor=j
			}
		}
	}
	fmt.Println(a)
}