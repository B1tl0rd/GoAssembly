package main

import (
	"time"

	"github.com/AndrusGerman/GoAssembly"
)

func RutaVariables() *GoAssembly.Page {
	var App GoAssembly.Page
	App.Template = `
	<button @rt="index">Go to Index</button>
	<center>
		<h1>Pagina Variables</h1>
		<h3>Name contiene {{ name }} </h3>
		<h3>Nivel 0-> {{ nnA1 }}</h3>
		<h3>Nivel 1-> {{ nnB2 }}</h3>
		<h3>Nivel 2-> {{ nnC3 }}</h3>
	</center>
	`
	//Se declaran la variables
	App.Data = func() {
		//Variables tipo 0 -> Se eliminan al cambiar de ruta
		App.Var.Set("name", "Andrus")
		App.Var.Set("nnA1", 0)
		//Variables tipo 1 -> Se mantienen en toda la seccion
		App.Var.SetT("nnB2", 0, false)
		//Variables tipo 2 -> Se mantienen en el LocalStorea del navegador -> Casi permanente
		App.Var.SetT("nnC3", 0, true)
	}
	//Script se ejecuta al preparar los datos y el Template
	App.Script = func() {
		for {
			time.Sleep(1 * time.Second)
			println("Sumando")
			a := App.Var.GetInt("nnA1") + 1
			b := App.Var.GetInt("nnB2") + 1
			c := App.Var.GetInt("nnC3") + 1
			App.Var.Set("nnA1", a)
			App.Var.SetT("nnB2", b, false)
			App.Var.SetT("nnC3", c, true)
		}
	}
	return &App
}
