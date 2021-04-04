package main

func main() {
	descriptor := [3][2]string{
		{"penguin", "cool"},
		{"hoarfroster", "handsome"},
		{"passionpenguin", "unknown"},
	}
	for i := 0; i < len(descriptor); i++ {
		for j := 0; j < len(descriptor[i]); j++ {
			println(descriptor[i][j])
		}
	}
}
