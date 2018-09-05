package main

import (
	"syscall/js"

	"github.com/AndrusGerman/GoAssembly"
)

func RutaRequest() *GoAssembly.Page {
	var App GoAssembly.Page
	//Axios := GoAssembly.NewAxios(&App)
	Doc := GoAssembly.Document(&App)
	App.Template = `
	<center>
		<h1>Request Page</h1>
		<button @e="get.click" >Get</button>
	</center>
	`
	App.Methods = func() {
		App.Method["get"] = func(a []js.Value) {
			//b := Axios.Get("http://www.facebook.com")
			Doc.Alert("No Avalible :(")
		}
	}

	//Prepara el index
	App.Prepare = func() {
		App.Title("Request")
	}
	return &App
}
