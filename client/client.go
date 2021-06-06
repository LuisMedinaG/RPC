package main

import (
	"fmt"
	"net/rpc"
)

type Alumno struct {
	Nombre       string
	Materia      string
	Calificacion float64
}

func main() {
	c, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	var opc int64
	for {
		fmt.Println("\n\t Menú Cliente\n")
		fmt.Println("1) Agregar calificación de una materia")
		fmt.Println("2) Mostrar promedio de un alumno")
		fmt.Println("3) Mostrar el promedio general")
		fmt.Println("4) Mostrar el promedio de una materia")
		fmt.Println("0) SALIR")
		fmt.Print("Opción: ")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			/*
				Pedir el nombre del alumno, materia y su correspondiente calificación, para posteriormente invocar por RPC la función para Agregar la calificación de un alumno por materia.
			*/
			fmt.Println("\n\t Agregar calificación\n")

			var nombre string
			var materia string
			var calificacion float64
			var reply bool

			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)
			fmt.Print("Calificación: ")
			fmt.Scanln(&calificacion)

			a := Alumno{nombre, materia, calificacion}
			err = c.Call("API.AgregarCalificacion", a, &reply)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("INFO: Calificacion agregada exitosamente")
		case 2:
			fmt.Println("\n\t Mostrar promedio de un alumno\n")

			var nombre string
			var result float64
			fmt.Print("-> Nombre : ")
			fmt.Scanln(&nombre)

			err = c.Call("API.MostrarPromedioAlumno", nombre, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio alumno ", nombre, ":", result)
			}
		case 3:
			fmt.Println("\n\t Mostrar promedio general\n")

			var result float64
			err = c.Call("API.MostrarPromedioGeneral", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio general:", result)
			}
		case 4:
			fmt.Println("\n\t Mostrar promedio de una materia\n")

			var materia string
			var result float64
			fmt.Print("-> Materia : ")
			fmt.Scanln(&materia)

			err = c.Call("API.MostrarPromedioMateria", materia, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio de la materia", materia, ":", result)
			}
		case 0:
			fmt.Println("\n\t Saliendo del programa . . .\n")
			return
		default:
			fmt.Println("\n\t Error, intenta de nuevo. . .\n")
		}
	}
}
