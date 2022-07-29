package models

import "fmt"

//paquete de modelos de calculadora
//struct Cohete encargado de llegar a la luna
type Cohete struct {
	Galones_gasolina int
	Tripulantes      int
	Mision           string
	Estado           string
	Motor            string
}

//funcion cargar gasolina
func (myCohete *Cohete) CargarCombustible(cantidadCombustible int) {
	myCohete.Galones_gasolina = myCohete.Galones_gasolina + cantidadCombustible
}

//funcion ingresar pasajeros
func (myCohete *Cohete) IngresarPasajeros(cantidadPasajeros int) {
	myCohete.Tripulantes = myCohete.Tripulantes + cantidadPasajeros
}

//funcion para asignar un nombre a la mision del cohete
func (myCohete *Cohete) NombreMision(nombreMision string) {
	myCohete.Mision = nombreMision
}

//funcion para cambiar el estado de la mision
func (myCohete *Cohete) EstadoMision(estadoMision string) {
	myCohete.Estado = estadoMision
}

//funcion para despegar
func (myCohete *Cohete) DespegarCohete(activarMotor string) {
	myCohete.Motor = activarMotor
}

//Salida del struct
func (myCohete Cohete) String() string {
	//si los tripulantes estan completos envia mensaje de Ok!
	if myCohete.Tripulantes == 5 {
		MsjTripulantes := "Tripulacion Completa!"
		return fmt.Sprintf("Gasolina: %d Galones de 100, %d Tripulantes de 5, %s!, Mision: %s, Estado: %s", myCohete.Galones_gasolina, myCohete.Tripulantes, MsjTripulantes, myCohete.Mision, myCohete.Estado)
	}
	//se agrega logica para trabajar la salida de la funcion
	//recordar %d para int y %s string
	return fmt.Sprintf("Gasolina: %d Galones de 100, %d Tripulantes de 5, Mision: %s, Estado: %s", myCohete.Galones_gasolina, myCohete.Tripulantes, myCohete.Mision, myCohete.Estado)
}
