package main

import (
	"fmt"
	"math"
)

func main() {
	var i, j = 3, 2
	i, j = j, i
	fmt.Println("i =", i, ",", "j = ", j)
	fmt.Println(address() == address())

	fmt.Printf("%08b\n", i)
	fmt.Printf("%08b\n", j)
	fmt.Println(math.NaN())

	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)

	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
	fmt.Println(1i * 1i)

	if err := testRecover(); err != nil {
		fmt.Println(err)
	}
}

// address as we all known, the local var's address is safe
func address() *int {
	v := 1
	return &v
}

func testRecover() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()

	panic("test")
	return nil
}
