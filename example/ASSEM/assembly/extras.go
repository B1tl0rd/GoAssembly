package main

import (
	"time"

	"github.com/AndrusGerman/GoAssembly"
)

func RutaExtra() *GoAssembly.Page {
	var App GoAssembly.Page
	App.Template = `
	<button @rt="index">Go to Index</button>
	<center>
		<h1>OnApp</h1>
		<h1>{{ numero }}</h1>
	</center>
	`
	//EL script que iniciar despues de que todo este preparado
	App.Script = func() {
		for {
			time.Sleep(1 * time.Second)
			n := App.Var.GetInt("numero") + 1
			App.Var.Set("numero", n)
		}
	}
	//Inicia antes de mostrar el template
	App.Prepare = func() {
		App.Title("Mi App")
	}
	//Declara tus variables
	App.Data = func() {
		App.Var.Set("numero", 0)
	}
	return &App
}
