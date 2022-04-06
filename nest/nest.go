package nest

func F() {
	var gopher int // want "pattern"
	print(gopher)  // want "identifier is gopher"
}

func F2() {
	var gopher int // want "pattern"
	print(gopher)  // want "identifier is gopher"
}

func F3() {
	var gopher int // want "pattern"
	print(gopher)  // want "identifier is gopher"
}
