package main

import (
	"fmt"
)

func main() {
	fibo := fiboFac()

    for i := 0; i < 10; i++ {
        fmt.Println(fibo())
    }
}

func fiboFac() func() int {
	a, b, tmp := 0, 1, 0
    return func() int {
		tmp = a + b
		a = b
		b = tmp

        return a
    }
}
