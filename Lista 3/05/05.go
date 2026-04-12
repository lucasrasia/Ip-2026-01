package main

import "fmt"

type pessoa struct {
	idade  int
	altura float64
	peso   float64
}

func main() {
	var ind []pessoa
	var i int
	var n string
	var h, p float64
	for {
		fmt.Print("Idade: ")
		fmt.Scan(&i)
		fmt.Print("Altura: ")
		fmt.Scan(&h)
		fmt.Print("Peso: ")
		fmt.Scan(&p)
		pessoaData := pessoa{
			idade:  i,
			altura: h,
			peso:   p,
		}
		ind = append(ind, pessoaData)
		fmt.Print("Deseja continuar? ('s' ou 'n') ")
		fmt.Scan(&n)
		if n == "n" {
			break
		}
	}
	var a int
	for i := range ind {
		if ind[i].idade > 50 {
			a = a + 1
		}
	}
	fmt.Printf("\nhá %v pessoas com mais de 50 anos\n", a)
	var b int     //qntd de pessoas entre 10 e 20
	var c float64 //somatório das alturas
	for i := range ind {
		if ind[i].idade > 10 && ind[i].idade < 20 {
			b = b + 1
			c = c + ind[i].altura
		}
	}
	var m float64
	if b == 0 {
		m = 0
	} else {
		m = c / float64(b)
	}
	fmt.Printf("a média da altura das pessoas entre com idade entre 10 e 20 anos é %v\n", m)
	var d float64 //quantidade de pessoas acima de 40kg
	var e float64 //total de pessoas
	for i := range ind {
		e = e + 1
		if ind[i].peso > 40 {
			d = d + 1
		}
	}
	p = d / (e)
	fmt.Printf("porcentagem de pessoas acima de 40kg: %.2f\n", p)
}
