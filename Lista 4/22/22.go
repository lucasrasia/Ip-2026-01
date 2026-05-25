package main
import (
	"fmt"
	"os"
	"os/exec"
)	
func limparTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main(){
	contas:=make(map[int]float64)
	// banco de contas com map
	contas[1]=500
	contas[2]=4327
	contas[3]=647.8
	contas[4]=100000.5
	contas[5]=5.7
	contas[6]=6767.6
	contas[7]=500
	contas[8]=500
	contas[9]=4834
	contas[10]=1000
	for{
		var id int
		fmt.Println(" Sistema bancário telecurso 2000")
		fmt.Println("================================")
		fmt.Println("[1] Efetuar depósito")
		fmt.Println("[2] Efetuar saque")
		fmt.Println("[3] Consultar o ativo bancário")
		fmt.Println("[4] Fechar sistema")
		var op int
		fmt.Print("O que deseja fazer? ")
		fmt.Scan(&op)
		saldo, existe := contas[id] //saldo=contas[id]     existe -> bool (tem uma chave=id -> true)
		if op==1{
			var dep float64
			limparTerminal()
			fmt.Print("Digite o id da conta: ")
			fmt.Scan(&id)
			if existe{
				fmt.Printf("Digíte o valor do depósito da conta %v (saldo atual: %v): ", id, saldo )
				fmt.Scan(&dep)
				contas[id] = saldo + dep
				limparTerminal()
				fmt.Println("Depósito efetuado com sucesso")
				continue
			} else {
				limparTerminal()
				fmt.Printf("Desculpe não existe uma conta com id: %v \n\n", id)
				continue
			}
		}
		if op==2{
			var saq float64
			limparTerminal()
			fmt.Print("Digite o id da conta: ")
			fmt.Scan(&id)
			saldo, existe := contas[id] 
			if existe{
				fmt.Printf("Digíte o valor do saque da conta %v (saldo atual: %v): ", id, saldo )
				fmt.Scan(&saq)
				if saq>contas[id]{
					limparTerminal()
					fmt.Println("Desculpe, saque maior que o saldo\n")
					continue
				}
				contas[id] = saldo - saq
			} else {
				limparTerminal()
				fmt.Printf("Desculpe não existe uma conta com id: %v \n\n", id)
				continue
		}
	}
}
}