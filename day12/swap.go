package main

func swap(x, y string) {
	x, y = y, x
}

func swapRef(x, y *string) {
	*x, *y = *y, *x
}
