package main

func main() {

	a := "amma"
	b := len(a) - 1
	c := 0
	for i := 0; i < len(a); i++ {
		if a[i] == a[b] {
			c++
			b--
		}
	}
	if c == len(a) {
		println("palindrom soz")

	} else {
		println("palindrom emas")
	}
}
