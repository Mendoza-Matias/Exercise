package main

import "fmt"

//creo una estructura para representar la tarea
type Tarea struct {
	id     int
	nombre string
}

var id int = 0 /* creo una variable global */

var lista []Tarea /* lista donde voy a almacenar las tareas */

func main() {
	agregarTarea("Comprar pan")
	agregarTarea("Estudiar")
	listarTareas()
}
func agregarTarea(nombre string) {
	aumentar(&id) /* le paso la referencia en la memoria */
	var t Tarea = Tarea{
		id:     id,
		nombre: nombre,
	}
	lista = append(lista, t)
}
func listarTareas() {
	for _, v := range lista {
		fmt.Println(v)
	}
}
func aumentar(id *int) { /* espero un parametro de tipo puntero */
	*id += 1 /* modifico el valor de la variable */
}
