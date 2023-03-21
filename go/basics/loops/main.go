package main

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		println(i)
	}

	// while loop
	i := 0
	for i < 5 {
		println(i)
		i++
	}

	// range loop
	s := []int{1, 2, 3}
	for i, v := range s {
		println(i, v)
	}

	// break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		println(i)
	}

	// continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		println(i)
	}

}
