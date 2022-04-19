package main

import "fmt"

func main() {
	primeNos := generatePrimes(3, 100)
	fmt.Println(primeNos)
}

func generatePrimes(start, end int) []int {
	primeNos := []int{}
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primeNos = append(primeNos, no)
		}
	}
	return primeNos
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
