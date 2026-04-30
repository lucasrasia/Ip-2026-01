package main

import "fmt"

func main() {
	var A []int
	var a int
	for i := 1; i <= 10; i++ {
		fmt.Printf("digito %v: ", i)
		fmt.Scan(&a)
		A = append(A, a)
	}
	soma:=make(map[int]int)
	for _, i:=range A{	
		soma[i]++
	}
	for i, qnt :=range soma{
		if qnt>1{
			fmt.Println(i, "aparece", qnt, "vezes")
		}
	}
}
