package main

func main() {

	println("maps")

	m1 := make(map[string]int)

	m1["a"] = 1
	m1["b"] = 2
	m1["c"] = 3

	for clave, valor := range m1 {
		println(clave, valor)
	}

}
