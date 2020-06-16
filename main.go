package main

func main() {
	var message string = "Hola Mundo, GO"
	// es la forma mas comun de declarar
	message2 := "Mensaje de prueba"

	/*
	  Imprimir variables
	*/
	println(message)
	println(message2)
	privada()
	Publica()
	testDefer()
}

func privada() {
	println("Funcion privada")
}

func Publica() {
	println("Funcion Publica")
}

func testDefer() {
	defer println("Mundo")
	println("Hola")
}
