package main

import "fmt"

type ErrNegative float64

type ErrZero int

func (e ErrNegative) Error() string {
	return "Entered value is negative."
}
func (e ErrZero) Error() string {
	return "Entered value is zero."
}

func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, ErrNegative(n)
	} else if n == 0 {
		return 0, ErrZero(n)
	} else {
		return n * n, nil
	}
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
	fmt.Println(Sqrt(0))
}
