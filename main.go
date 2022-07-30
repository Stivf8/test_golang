package main

import (
	"fmt"
	lunch "test/models"
	basicFunctions "test/packages"
)

func main() {

	res1 := basicFunctions.Sumar(69, 1)

	fmt.Println("desde funcion Sumar!: ", res1)

	//struct de cohete ejecucion
	//instanciar
	//inicio del cohete
	CoheteMark := lunch.Cohete{Galones_gasolina: 0, Tripulantes: 0, Mision: "Sin Asignar", Estado: "Pendiente"}
	fmt.Println("==============================Cohete en plataforma de despegue!==============================")
	fmt.Println("Cohete: ", CoheteMark)
	//se carga la gasolina y los tripulantes
	fmt.Println("==============================Se actualiza estado del Cohete!==============================")
	CoheteMark.CargarCombustible(100)
	CoheteMark.IngresarPasajeros(5)
	CoheteMark.EstadoMision("En proceso")
	fmt.Println("Cohete: ", CoheteMark)
	//se asigna mision al cohete y esta listo para despegar
	CoheteMark.NombreMision("Ir a Marte")
	CoheteMark.EstadoMision("Listo para despegar!")
	fmt.Println("==============================Se actualiza estado del Cohete!==============================")
	fmt.Println("Cohete: ", CoheteMark)
	//Despega el Cohete
	CoheteMark.DespegarCohete("Iniciado")
	CoheteMark.EstadoMision("Despegando")
	//devolvemos mision
	msj := CoheteMark.GetMision()
	println(msj)

}
