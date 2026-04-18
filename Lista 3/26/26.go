package main
import "fmt"
func fat(v int)int{
	if v==1 || v==0{
		return 1
	}
	return v*fat(v-1)
}
func main(){
	n:=100.0
	var r float64
	for i:=0; i<20; i++{
		r=r+n/float64(fat(i))
		n=n-1
	}
	fmt.Println(r)
}