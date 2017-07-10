package main

func main() {
	x := 100

	// If
	if x > 0 {
		println("x")
	} else if x < 0 {
		println("-x")
	} else {
		println("0")
	}

	// Switch
	switch {
	case x > 0:
		println("x")
	case x < 0:
		println("-x")
	default:
		println("0")
	}

	// For
	for i := 0; i < 5; i++ {
		println(i)
	}

	for i := 4; i >= 0; i-- {
		println(i)
	}

	// 相当于 while(x<5) {...}
	for x < 5 {
		println(x)
		x++
	}

	x = 4
	// 项目于while(true) {...}
	for {
		println(x)
		x--
		if x < 0 {
			break
		}
	}

	b := []int{100, 101, 102}

	for i, n := range b {
		println(i, ":", n)
	}
}
