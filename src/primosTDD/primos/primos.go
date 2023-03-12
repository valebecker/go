package primos

import "fmt"

func EsPrimo(x int) bool {
	if x < 2 {
		return false
	}

	for n := 2; n < x/2+1; n++ {
		if x%n == 0 {
			return false
		}
	}
	return true
}

func ListarPrimos(n int) {
	contador := 1
	i := 0
	var primo_grande int
	for contador <= n {
		i++
		if EsPrimo(i) {
			primo_grande = i
			contador++
		}
	}

	fmt.Println(primo_grande)
}
