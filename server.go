package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"

	alumno "./common"
)

var materias = make(map[string]map[string]float64)

type API struct

func (this *API) AgregarCalificacion(stu alumno.Alumno, reply *bool) error {

	materia := make(map[string]float64)
	materia[stu.Materia] = stu.Calificacion

	listSize := len(alumnos)

	if alumno, err := alumnos[stu.Nombre]; err {
		alumno[stu.Materia] = stu.Calificacion
		if len(alumno) > listSize {
			*reply = true
		}
	} else {
		// alumnos[stu.Nombre] = materia
		if len(alumnos) > listSize {
			*reply = true
		}
		//fmt.Println("Número de alumnos:", len(alumnos))
	}

	return nil
}

func (this *API) MostrarPromedioAlumno(nombre string, reply *float64) error {
	flag := false
	// comprobemos que la lista no esté vacía
	if len(alumnos) > 0 {
		// comprobemos que haya registros de ese alumno
		for valor := range alumnos {
			if valor == nombre {
				flag = true
			}
		}
		if flag {
			// obtenemos el promedio
			var sumatoria float64
			for _, cal := range alumnos[nombre] {
				sumatoria = sumatoria + cal
			}
			promedio := sumatoria / float64(len(alumnos[nombre]))
			*reply = promedio
		} else {
			return errors.New("No existe ese alumno")
		}
	} else {
		return errors.New("No hay elementos registrados")
	}

	return nil
}

func (this *API) MostrarPromedioGeneral(promedioGeneral float64, reply *float64) error {

	if len(alumnos) > 0 { // si la lista no está vacía
		var prom float64
		numAlumnos := float64(len(alumnos))

		for valor := range alumnos {
			var sumatoria float64
			for _, cal := range alumnos[valor] {
				sumatoria = sumatoria + cal
			}
			prom = sumatoria / float64(len(alumnos[valor]))
			promedioGeneral = promedioGeneral + prom
		}
		promedioGeneral = promedioGeneral / numAlumnos
		*reply = promedioGeneral
	} else {
		return errors.New("No hay elementos registrados")
	}

	return nil
}

func (this *API) MostrarPromedioMateria(materia string, reply *float64) error {

	if len(alumnos) > 0 {
		// buscaremos la materia en cada alumno registrado
		flag := false
		var sumatoria float64
		contador := 0
		for valor := range alumnos {
			for mat, cal := range alumnos[valor] {
				if mat == materia {
					flag = true
				}
				if flag {
					sumatoria = sumatoria + cal
					contador = contador + 1
				}
			}
		}

		promedio := sumatoria / float64(contador)
		*reply = promedio

	} else {
		return errors.New("No hay elementos registrados")
	}

	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error registrando API", err)
	}

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error arrancando servidor", err)
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("Error aceptando conexion", err)
			continue
		}

		go rpc.ServeConn(c)
	}
}
