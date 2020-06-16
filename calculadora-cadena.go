package main

type datos struct {
	numero1   int
	numero2   int
	operacion string
}

type calculadora interface {
	calcular(*datos)
	setNext(calculadora)
}

type sumar struct {
	next calculadora
}

type restar struct {
	next calculadora
}

type multiplicar struct {
	next calculadora
}

func (s *sumar) setNext(next calculadora) {
	s.next = next
}

func (r *restar) setNext(next calculadora) {
	r.next = next
}

func (m *multiplicar) setNext(next calculadora) {
	m.next = next
}

func (s *sumar) calcular(d *datos) {
	if d.operacion == "+" {
		println("La suma de los numeros es: ", d.numero1+d.numero2)
		return
	} else {
		s.next.calcular(d)
	}
}

func (r *restar) calcular(d *datos) {
	if d.operacion == "-" {
		println("La resta de los numeros es: ", d.numero1-d.numero2)
		return
	} else {
		r.next.calcular(d)
	}
}

func (m *multiplicar) calcular(d *datos) {
	if d.operacion == "*" {
		println("La multiplicacion de los numeros es: ", d.numero1*d.numero2)
		return
	} else {
		println("Operacion no soportada")
	}
}

func main() {
	sumar := &sumar{}
	restar := &restar{}
	multiplicar := &multiplicar{}

	sumar.setNext(restar)
	restar.setNext(multiplicar)

	datos := &datos{operacion: "+", numero1: 6, numero2: 4}

	sumar.calcular(datos)
}
