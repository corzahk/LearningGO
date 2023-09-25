package main

import "fmt"

// variables  y constantes

const MiConstante = "El valor de mi constante"

func main() {
	//declaracion por inferencia
	var nombre string = "Cesar"
	fmt.Println(nombre)
	//declaracion rápida o corta
	nombre2 := "Juan"
	fmt.Println(nombre2)
	fmt.Printf("El valor de MiConstante es %s", MiConstante)
}
