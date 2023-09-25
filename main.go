package main

import "fmt"

//condicionales
func main() {

	/*
		operadores de comparacion
		x == y ||es igual x igual a y?
		x != y ||es x diferente de y?
		x < y  ||es x menor que y
		x > y  || es x mayor que y
		x >= y  || es x mayor o igual que y?
	*/
	edad := 25
	if edad >= 17 {
		fmt.Println("Es mayor de edad")
	} else {
		fmt.Println("Es menor de edad")
	}

	//declarar variable en una condicion
	if variable := 2; variable == 2 {
		fmt.Println("----------------")
		fmt.Println("variables es igual a 2")
	}
}

/*
//punteros
// & sirve para poder acceder al valor de la memoria de esa variable
func main() {
	color := "rojo"
	fmt.Println(&color)
}
*/
/*
import (
	"fmt"
	"reflect"
)
// reclect y TypeOf
// reflec.typeOf sirve para ver que tipo de dato es  string, entero, float etc...
func main() {
	var string1 string = "algo ñangú"
	fmt.Println(reflect.TypeOf(string1))
}
*/
/*
//tipos de datos
func main() {
	var string string = "texto"
	fmt.Println(string)
	fmt.Println("------------------------------------")
	textoGrande := `Eu aute et reprehenderit labore voluptate. Irure qui fugiat culpa irure ipsum nostrud deserunt amet officia Lorem fugiat dolor ullamco. Esse do do excepteur ad laboris labore fugiat cillum qui ex ullamco exercitation. Sit dolor laborum non et. Reprehenderit dolore cupidatat minim nostrud nostrud fugiat.`
	fmt.Println(textoGrande)
	fmt.Println("------------------------------------")
	var estado bool = false
	fmt.Println(estado)
	fmt.Println("------------------------------------")
	var flotante32 float32 = 34.66
	fmt.Println(flotante32)
	fmt.Println("------------------------------------")
	var flotante64 float64 = 54.65
	fmt.Println(flotante64)
	fmt.Println("------------------------------------")
	var entero int = 1234
	fmt.Println(entero)
	fmt.Println("------------------------------------")
	var entero_int8 int8 = 123 //-128 a 127
	fmt.Println(entero_int8)
	fmt.Println("------------------------------------")
	var entero_int16 int16 = 123 //-2*15 a 2 *15 -1
	fmt.Println(entero_int16)
	fmt.Println("------------------------------------")
	var entero_int32 int32 = 78910 //-2*63a 2 *63 -1
	fmt.Println(entero_int32)
	fmt.Println("------------------------------------")
	var entero_int64 int64 = 1223 //0 a 255
	fmt.Println(entero_int64)
	fmt.Println("------------------------------------")
	var entero_uint8 uint8 = 122 //0 a 2*8 -1
	fmt.Println(entero_uint8)
	fmt.Println("------------------------------------")
	var entero_uint16 uint16 = 42343 //0 a 2*16 -1
	fmt.Println(entero_uint16)
	fmt.Println("------------------------------------")
	var entero_uint32 uint32 = 312312312 //0 a 2*32 -1
	fmt.Println(entero_uint32)
	fmt.Println("------------------------------------")
	var entero_uint64 uint64 = 423535343343 //0 a 2*64 -1
	fmt.Println(entero_uint64)
	fmt.Println("------------------------------------")
}
*/
/*
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
*/
