package main

import "fmt"

type TemporaryEmpleado struct {
	Persona
	Empleado
	taxRate int
}

func (ftEmpleado TemporaryEmpleado) getMessage() string {
	return "Empleado Temporal"
}

type Persona struct {
	edad int
}

type Empleado struct {
	id       int
	name     string
	vacation bool
}

type FullTimeEmpleado struct {
	Persona
	Empleado
	endDate string
}

func (ftEmpleado FullTimeEmpleado) getMessage() string {
	return "Empleado full time"
}

type ImprimirInfo interface {
	getMessage() string
}

func getMessage(p ImprimirInfo) {
	fmt.Println(p.getMessage())
}

func NuevoEmpleado(id int, name string, vacation bool) *Empleado {
	return &Empleado{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

// pointer receivers function,
// son métodos por referencia es decir modifican el valor que es enviado
func (e *Empleado) SetId(id int) {
	e.id = id
}

func (e *Empleado) SetName(name string) {
	e.name = name
}

func (e *Empleado) GetID() int {
	return e.id
}

func (e *Empleado) GetName() string {
	return e.name
}

func main() {
	e := Empleado{}
	fmt.Println(e)
	e.id = 1
	e.name = "Nombre"
	fmt.Println(e)
	e.SetId(5)
	e.SetName("Aco")
	fmt.Println(e.GetID(), e.GetName())
	// constructor
	e2 := Empleado{
		id:       1,
		name:     "Un Nombre",
		vacation: true,
	}
	fmt.Println(e2)

	// 3
	e3 := new(Empleado)
	fmt.Println(*e3)

	e4 := NuevoEmpleado(4, "otro nombre", false)
	fmt.Println(*e4)

	// el polimorfismo no funciona en go, aunque no es necesario
	ftEmpleado := FullTimeEmpleado{}
	ftEmpleado.name = "Aco"
	ftEmpleado.edad = 28
	ftEmpleado.id = 7
	ftEmpleado.vacation = true

	fmt.Println(ftEmpleado)

	tEmpleado := TemporaryEmpleado{}
	getMessage(ftEmpleado)
	getMessage(tEmpleado)
}
