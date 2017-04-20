package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/koshalt/stringutil"
)

var e, f, g bool

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("hello, world", time.Now())
	fmt.Println(stringutil.Reverse(" !oG ,olleH"), rand.Intn(10))

	fmt.Println(add(5, 6))
	fmt.Println(mul(5, 6))

	a, b := swap("me", "swap")
	fmt.Println(a, b)

	sA, sB := split(10)
	fmt.Println(sA, sB)

	var h int
	fmt.Println(e, f, g, h)

	fmt.Println("Sum to 10:", sumUpto(10))
	fmt.Println("Fib sum to 1000:", fibSum(1000))

	fmt.Println("switching")
	switchGo()
}

func add(a int, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(toSplit int) (a, b int) {
	a = toSplit / 2
	b = toSplit - toSplit/2
	return
}

func sumUpto(upto int) int {
	sum := 0
	for i := 0; i <= upto; i++ {
		sum += i
	}
	return sum
}

func fibSum(upto int) (sum int) {
	sum = 1
	for sum < upto {
		sum += sum
	}
	return
}

func switchGo() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("%s", os)
	}
}
