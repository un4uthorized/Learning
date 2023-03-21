package main

func main() {
	// if statement
	if 7%2 == 0 {
		println("7 is even")
	} else {
		println("7 is odd")
	}

	// if with a short statement
	if num := 9; num < 0 {
		println(num, "is negative")
	} else if num < 10 {
		println(num, "has 1 digit")
	} else {
		println(num, "has multiple digits")
	}

	// switch statement
	switch os := "darwin"; os {
	case "darwin":
		println("OS X.")
	case "linux":
		println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		println("%s.", os)
	}
}
