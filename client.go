package main

import (
	"fmt"
	"net/rpc"

	alumno "./common"
)

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	var opc int64
	for {
		fmt.Println("\n\n\t Menú Cliente\n")
		fmt.Println("1) Agregar calificación de una materia")
		fmt.Println("2) Mostrar promedio de un alumno")
		fmt.Println("3) Mostrar el promedio general")
		fmt.Println("4) Mostrar el promedio de una materia")
		fmt.Println("0) SALIR")
		fmt.Print("-> Opción: ")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			/*
				Pedir el nombre del alumno, materia y su correspondiente calificación, para posteriormente invocar por RPC la función para Agregar la calificación de un alumno por materia.
			*/
			fmt.Println("\n\t Agregar calificación\n")

			// hay que pedir nombre del alumno, materia y calificación
			var name string
			var course string
			var calif float64
			result := false

			fmt.Print("Nombre: ")
			fmt.Scanln(&name)
			fmt.Print("Materia: ")
			fmt.Scanln(&course)
			fmt.Print("Calificación: ")
			fmt.Scanln(&calif)

			a := alumno.Alumno{name, course, calif}
			err = c.Call("API.AgregarCalificacion", a, &result)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("INFO: Calificacion agregada exitosamente")
		case 2:
			/*
				Pedir nombre del alumno, invocar por RPC la función para obtener e imprimir el promedio del alumno.
			*/
			fmt.Println("\n\t Mostrar promedio de un alumno\n")

			var name string
			var result float64
			fmt.Print("-> Nombre : ")
			fmt.Scanln(&name)

			err = c.Call("API.MostrarPromedioAlumno", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("API.MostrarPromedioAlumno =>", "Promedio:", result)
			}
		case 3:
			/*
				Invocar por RPC la función Obtener el promedio de todos los alumnos e imprimir el promedio general.
			*/
			fmt.Println("\n\t Mostrar promedio general\n")

			var result float64
			err = c.Call("API.MostrarPromedioGeneral", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("API.MostrarPromedioGeneral =>", "Promedio general:", result)
			}
		case 4:
			/*
				Invocar por RPC la función Obtener el promedio de todos los alumnos e imprimir el promedio general.
			*/
			fmt.Println("\n\t Mostrar promedio de una materia\n")

			var materia string
			var result float64
			fmt.Print("-> Materia : ")
			fmt.Scanln(&materia)

			err = c.Call("API.MostrarPromedioMateria", materia, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("API.MostrarPromedioMateria =>", "Promedio de la materia:", result)
			}
		case 0:
			fmt.Println("\n\t Saliendo del programa . . .\n")
			return
		default:
			fmt.Println("\n\t [!] Error, intenta de nuevo. . .\n")
		}
	}
}

func main() {
	client()
}
