package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

var c, python, java bool // variables del paquete

var i, j int = 1, 2 // inicializado de variables

// delaramos varias variables de diferentes tipos de esta manera
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("Hola mundo!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var i int // variable de la función
	k := 3    // declara e inicializa, esta forma solo puede usarse dentro de funciones
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)

	var i1 int
	var f1 float64
	var b1 bool
	var s1 string
	fmt.Printf("%v %v %v %q\n", i1, f1, b1, s1)

	//For
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//Como el while de java
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	//switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	//Evitamos los largos if-then-else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//Estructuras
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v)

	//Estructuras con punteros
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)

	//Arrays
	var a1 [2]string
	a1[0] = "Hello"
	a1[1] = "World"
	fmt.Println(a1[0], a1[1])
	fmt.Println(a1)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	//Otros usos
	s2 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	s2 = s2[1:4]
	fmt.Println(s)

	s2 = s2[:2]
	fmt.Println(s)

	s2 = s2[1:]
	fmt.Println(s)

	//largo y capacidad de un slice
	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3)

	// Slice the slice to give it zero length.
	s3 = s3[:0]
	printSlice(s3)

	// Extend its length.
	s3 = s3[:4]
	printSlice(s3)

	// Drop its first two values.
	s3 = s3[2:]
	printSlice(s3)

}

func add(x int, y int) int {
	return x + y
}

// Las funciones pueden devolver varios resultados
func swap(x, y string) (string, string) {
	return y, x
}

// Los valores de retorno pueden ser denominados
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // esto se llama naked return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func punteros() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func ticTacToe() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func concatenar() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}
