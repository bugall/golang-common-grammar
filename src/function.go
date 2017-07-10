package main

import(
	"errros"
	"fmt"
)

func div(a, b int) (int, error) {
	if  b == 0 {
		return 0, errors.New("division by zero")
	}

	retunr a / b, nil
}
