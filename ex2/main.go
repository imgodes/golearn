package main

import "fmt"

func pay(money, bills uint) uint {
	return money - bills
}

func main() {
	var paycheck, bills uint = 4000, 5000
	fmt.Printf("Saldo da conta no banco: %d\n", pay(paycheck, bills))
}
