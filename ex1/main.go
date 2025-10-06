package main

import "fmt"

func pay(money, bills int) int {
	return money - bills
}

func main() {
	var paycheck, bills int = 4000, 5000
	fmt.Printf("Saldo da conta no banco: %d\n", pay(paycheck, bills))
}
