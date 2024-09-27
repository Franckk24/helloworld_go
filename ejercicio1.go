package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func imprmir() {
	fmt.Println("texto de la funcion imprimir")
}

func hola_string(s string) string {
	return "hola " + s
}

func sumar(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("a es igual a b")
	} else {
		fmt.Println("a no es igual a b")
	}
}

func arreglo() {
	var aprendices [5]string
	aprendices[0] = "leonardo"
	aprendices[1] = "juan sebastian"
	aprendices[2] = "franck"
	aprendices[3] = "juan sebastian"
	aprendices[4] = "jaider"

	fmt.Println(aprendices[0])
}

func tipos_de_datos() {
	var texto string = "fabian"
	var numero int = 1000
	var decimal float64 = 0.00001

	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))
}

func convert_stringtoboolean (){
	var palabra string = "true"
	boolean, err := strconv.ParseBool(palabra)
	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Println("este es el erro ",err )
}

func converbooltostring () {
	var palabra_bool = true
	strbool := strconv.FormatBool(palabra_bool)
	fmt.Println(strbool, reflect.TypeOf(strbool))
}










func variablesenunasolalínea () {
    var nombre, apellido string = "leonardo", "cruz"
    fmt.Println(nombre, apellido)
}

func asignarmúltiplesvariables () {
    var (
        nombre     string = "leonardo"
        edad       int    = 20
        aprendiz   bool   = true
    )
    fmt.Println("Nombre: ", nombre)
    fmt.Println("Edad: ", edad)
    fmt.Println("aprendiz: ", aprendiz)
}

func valores_por_defecto() {
    var palabra string
    var edad int
    var peso float64
    var estudiante bool
    fmt.Println("Nombre: ", palabra)
    fmt.Println("Edad: ", edad)
    fmt.Println("Peso: ", peso)
    fmt.Println("Estudiante: ", estudiante)
}

func usodeloperador() {
	nombre := "leonardo"
	edad := 20
	peso := 90
	estudiante := true 

	fmt.Printf("Nombre: ", nombre)
	fmt.Printf("Edad: ", edad)
	fmt.Printf("Peso: ", peso)
	fmt.Printf("Estudiante: ", estudiante)
}

var narcotraficante = "pablo escobar"

func definicióndevariableglobal () {
	cocaina := 10 
	fmt.Println("Cocaina: ", cocaina)
	fmt.Println("Narcotraficante: ", narcotraficante)
}

var variable1 = "este es el nivel 1"

func usodescopeenGo () {
	var variable2 = "este es el nivel 2"
	{
		var variable3 = "este es el nivel 3"
		fmt.Println(variable3)
	}
	fmt.Println(variable1)
	fmt.Println(variable2)
}


func sencillodelusodepunteros () {
	color := "blaco"
	fmt.Println(color, &color) 
}

func usodepunteros () {
	vehiculo1 := "verde"
	fmt.Println("el vehiculo1 es ", &vehiculo1)

	vehiculo2 := vehiculo1
	fmt.Println("el vehiculo2 es ", vehiculo2)

	vehiculo3 := &vehiculo1
	fmt.Println("el vehiculo3 es ", *vehiculo3)

	vehiculo1 = "gris" 

	fmt.Println("vehiculo1 es ", vehiculo1, &vehiculo1)
	fmt.Println("vehiculo2 es ", vehiculo2, &vehiculo2)
	fmt.Println("vehiculo3 es ", *vehiculo3, vehiculo3)
}


func equivalenciaEnPies(altura float32) float32 {
    altura = altura * 3.28
    return altura
}

var altura float32 = 1.70

func conversionEnPies(altura *float32) float32 {
    *altura = *altura - 0.10
    return *altura
}

func main() {
	// fmt.Println("texto de la funcion main")
	// imprmir()
	// fmt.Println(hola_string("fabian"))
	// fmt.Println(sumar(3, 5))
	// comparar_bool()
	// arreglo()
	// tipos_de_datos()
	// convert_stringtoboolean()
	// converbooltostring()
	fmt.Println("La altura es:", altura, "mts")
	fmt.Println("La altura es:", equivalenciaEnPies(altura), " pies")
	fmt.Println("La nueva altura es:", altura, "mts")
	
}
