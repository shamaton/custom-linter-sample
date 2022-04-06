package nest

func F() {
	var gopher int // want "pattern"
	print(gopher)  // want "identifier is gopher"
}
