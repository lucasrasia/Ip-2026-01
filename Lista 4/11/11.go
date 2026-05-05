package main

import (
	"fmt"
	"math"
)
func main(){
	var a[]float64
	var b float64
	for i:=1; i<=100; i++{
		fmt.Scan(&b)
		a=append(a, b)
	}
	var s float64
	j:=99
	for i:=0;i<=50;i++{
		s=s+math.Pow((a[i]-a[j]), 3)
		j-=1
	}
}