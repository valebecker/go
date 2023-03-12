package primos_test

import (
	"primosTDD/primos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeros(t *testing.T) {
	t.Parallel()
	b := primos.EsPrimo(0)
	assert.False(t, b, "Cero no es primo")
}
func TestUno(t *testing.T) {
	t.Parallel()
	b := primos.EsPrimo(1)
	assert.False(t, b, "1 no es primo")
}
func TestDos(t *testing.T) {
	t.Parallel()
	b := primos.EsPrimo(2)
	assert.True(t, b, "2 es primo")
}
func TestTres(t *testing.T) {
	t.Parallel()
	b := primos.EsPrimo(3)
	assert.True(t, b, "3 es primo")
}
func TestCuatro(t *testing.T) {
	t.Parallel()
	b := primos.EsPrimo(4)
	assert.False(t, b, "4 no es primo")
}
func TestCinco(t *testing.T) {
	t.Parallel()
	b := primos.EsPrimo(5)
	assert.True(t, b, "5 es primo")
}
func TestSiete(t *testing.T) {
	t.Parallel()
	b := primos.EsPrimo(7)
	assert.True(t, b, "7 es primo")
	b = primos.EsPrimo(89)
	assert.True(t, b, "89 es primo")
}
func TestNoPrimos(t *testing.T) {
	t.Parallel()
	b := primos.EsPrimo(10)
	assert.False(t, b, "10 No es primo")
	b = primos.EsPrimo(57)
	assert.False(t, b, "57 No es primo")
	b = primos.EsPrimo(121)
	assert.False(t, b, "121 No es primo")
	b = primos.EsPrimo(8633)
	assert.False(t, b, "8633 No es primo")
}

/*func TestListar(t *testing.T) {
	t.Parallel()
	primos.ListarPrimos(100000)
}*/
