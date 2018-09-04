package main

import (
	"syscall/js"

	"github.com/AndrusGerman/GoAssembly"
)

func RutaLogin() *GoAssembly.Page {
	var App GoAssembly.Page
	Document := GoAssembly.Document(&App)
	App.Template = `
	<button @rt="index">Go to Index</button>
	<center>
		<input placeholder="minimo 4" id="a">
		<input placeholder="minimo 4" id="b">
		<button @e="ini.click" >Iniciar</button>
		<p> {{ err }} </p>
		<p> {{ ok }} </p>
	</center>
	`
	App.Methods = func() {
		App.Method["ini"] = func(a []js.Value) {
			nombre := Document.Id("a").Value.String()
			B := Document.Id("b").Value.String()
			if len(nombre) < 4 || len(B) < 4 {
				App.Var.Set("ok", "")
				App.Var.Set("err", "Muy corto :/")
			} else {
				App.Var.Set("err", "")
				App.Var.Set("ok", "Permitido")
				Document.Alert("Gracias Por Iniciar")
				App.Var.SetT("user", nombre, false)
				GoAssembly.RunRoute("index")
			}
		}
	}
	//Se declaran la variables
	App.Data = func() {
		App.Var.Set("err", "")
		App.Var.Set("ok", "")
		App.Var.SetT("user", "", false)
	}
	//Prepara la pagina
	App.Prepare = func() {
		App.Title("Login Onfline")
	}
	return &App
}
