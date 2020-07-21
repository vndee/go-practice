package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"math"
	"math/rand"
	"strings"
	"time"
)

var t = 3

type Vertex struct {
	X int
	Y int
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	fmt.Printf("%d\n", sum(15))
	fmt.Println(time.Now())
	fmt.Println("My favorite number is", rand.Float64())
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Printf("Pi number is %.3g\n", math.Pi)
	fmt.Printf("Add: %d\n", add(1, 1))

	a, b := swap("abc", "bac")
	fmt.Printf("Swap abc - bac: (%s, %s)\n", a, b)

	fmt.Println(split(10))

	fmt.Println(t)

	vertex := Vertex{1, 2}
	fmt.Println(vertex.X)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	type Vertex struct {
		Lat, Long float64
	}
	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var m2 = map[string]Vertex{
		"Bell Labs": {
			40.45, 45.75,
		},
		"Google": {
			37.5, 55.5,
		},
	}
	fmt.Println(m2)

	m3 := make(map[string]int)
	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])
	m3["Answer"] = 43
	fmt.Println("The value:", m3["Answer"])
	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])
	v, ok := m3["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	wc.Test(WordCount)

	//hypot := func(x, y, float64) float64 {
	//	return math.Sqrt(x*x + y*y)
	//}
	//fmt.Println(hypot(5, 12))
	//fmt.Println(compute(hypot))
	//fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

func fibonacci() func(int) int {
	f1 := 0
	f2 := 1
	return func(x int) int {
		f := f1 + f2
		f1 = f2
		f2 = f
		return f
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func WordCount(s string) map[string]int {
	return map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sum(a int) int {
	sum := 0
	for i := 1; i <= a; i++ {
		sum += i
	}

	return sum
}

func add(a int, b int) int {
	return a + b
}

func swap(a string, b string) (string, string) {
	return a, b
}
