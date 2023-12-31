package main

import (
	"fmt"

	"github.com/IngeCamiloAriza/task-bots/business"
)

var uc = new(business.UseCase)

func main() {
	var option int
	fmt.Println("Hola master, que desea hacer ")
	fmt.Println("1. Consultar tarea para hoy \n2. Agregar actividades para la semana")
	fmt.Scanln(&option)

	optionBot(option)
}

func optionBot(option int) {
	if option == 1 {
		uc.SearchTaskDay()
	}
	if option == 2 {
		fmt.Printf("Todabia no esta diponible")
	}

}
