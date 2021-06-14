package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Alumno struct {
	Nombre       string
	Materia      string
	Calificacion float64
}

var materias = make(map[string]map[string]float64)

type API struct {
}

func (api *API) AgregarCalificacion(a Alumno, reply *bool) error {
	alumno := map[string]float64{
		a.Nombre: a.Calificacion,
	}

	if materia, ok := materias[a.Materia]; ok {
		materia[a.Nombre] = a.Calificacion
	} else {
		materias[a.Materia] = alumno
	}

	*reply = true
	return nil
}

func (api *API) MostrarPromedioAlumno(nombre string, reply *float64) error {
	var numMaterias float64
	var suma float64
	for _, materia := range materias {
		if calificacion, ok := materia[nombre]; ok {
			suma = suma + calificacion
			numMaterias += 1
		}
	}

	if numMaterias == 0 {
		return fmt.Errorf("ERROR: Alumno %s no existe", nombre)
	}

	*reply = suma / numMaterias

	return nil
}

func (api *API) MostrarPromedioGeneral(_ string, reply *float64) error {
	if len(materias) > 0 {
		var suma float64
		var numAlumnos float64
		for _, materia := range materias {
			for _, calificacion := range materia {
				suma += calificacion
				numAlumnos += 1
			}
		}
		*reply = suma / numAlumnos
	} else {
		return errors.New("no hay materias registradas")
	}

	return nil
}

func (api *API) MostrarPromedioMateria(titulo string, reply *float64) error {
	if materia, ok := materias[titulo]; ok {
		var suma float64
		var numAlumnos float64
		for _, calificacion := range materia {
			suma = suma + calificacion
			numAlumnos = numAlumnos + 1
		}

		*reply = suma / numAlumnos
	} else {
		return fmt.Errorf("ERROR: Materia %s no esta registrda", titulo)
	}

	return nil
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported.", http.StatusNotFound)
// 		return
// 	}

// 	fmt.Fprintf(w, "Hello!")
// }

func cargarHtml(a string) string {
	html, _ := ioutil.ReadFile(a)

	return string(html)
}

type Tarea struct {
	Nombre string
	Estado string
}

type AdminTareas struct {
	Tareas []Tarea
}

var misTareas AdminTareas

func (tareas *AdminTareas) Agregar(tarea Tarea) {
	tareas.Tareas = append(tareas.Tareas, tarea)
}

func (tareas *AdminTareas) String() string {
	var html string
	for _, tarea := range tareas.Tareas {
		html += "<tr>" +
			"<td>" + tarea.Nombre + "</td>" +
			"<td>" + tarea.Estado + "</td>" +
			"</tr>"
	}
	return html
}

func agregar(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("static/agregar.html"),
		)

	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		nombre := req.FormValue("nombre")
		materia := req.FormValue("materia")
		calif, _ := strconv.ParseFloat(req.FormValue("calif"), 64)

		alumno := Alumno{Nombre: nombre, Materia: materia, Calificacion: calif}
		fmt.Println(alumno)
		fmt.Println(materias)

		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("static/respuesta.html"),
			nombre,
			materia,
		)

		// AgregarCalificacion(alumno)
	}
}

func main() {
	// api := new(API)
	// err := rpc.Register(api)
	// if err != nil {
	// 	log.Fatal("Error registrando API", err)
	// }

	// fmt.Println("[INFO] API registrada")
	// ln, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	fmt.Println("Error arrancando servidor", err)
	// }

	// fmt.Println("[INFO] Arrancando servidor...")
	// for {
	// 	c, err := ln.Accept()
	// 	if err != nil {
	// 		fmt.Println("Error aceptando conexion", err)
	// 		continue
	// 	}

	// 	go rpc.ServeConn(c)
	// }

	// fileServer := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)

	http.HandleFunc("/agregar", agregar)

	fmt.Printf("Starting server at http://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
