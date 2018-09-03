package main

import "github.com/AndrusGerman/GoAssembly"
import "syscall/js"

func RutaEventos() *GoAssembly.Page {
	var App GoAssembly.Page
	Document := GoAssembly.Document(&App)

	App.Template = `
	<button @rt="index">Go to Index</button>
	<center>
		<h1>Pagina Eventos</h1>
		<button @e="hola.click">Imprime Nombre</button>
		<button @e="alerta.click">Alerta Nombre</button>
		<input @e="dato.input"/>
		<h1> {{ texto }}</h1>
	</center>
	`
	//Se pueden declara las variables
	App.Data = func() {
		App.Var.Set("texto", "")
	}
	//Se declaran los methods
	App.Methods = func() {
		App.Method["hola"] = func(dat []js.Value) {

			println("Hola Usuario")
		}
		App.Method["alerta"] = func(dat []js.Value) {
			Document.Alert("Hola Andrus")
		}
		App.Method["dato"] = func(dat []js.Value) {
			v := Document.Parse.GetValue(dat)
			//Document.Parse.GetId(dat) -> Id
			App.Var.Set("texto", v)
		}
	}
	return &App
}
