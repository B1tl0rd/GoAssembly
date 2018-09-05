# Welcome to GoAssembly!

Hi! I'm your first Assembly app in **GoAssembly**.

# GoAssembli
Go Assembly is one Framework for WebAssembly 


## Hello Word Code

```go
package main

import (
	"github.com/AndrusGerman/GoAssembly"
)

func main() {
	//Se declaran las rutas
	GoAssembly.Ruta["index"] = RutaIndex()
	//Run app
	GoAssembly.RunApp()
}


func RutaIndex() *GoAssembly.Page {
	// App
	var App GoAssembly.Page
	//Template
	App.Template = `
	<center>
		<h1>Hello Word</h1>
	</center>
	`
	//Prepare code
	App.Prepare = func() {
		App.Title("Mi First App")
	}
	return &App
}
```
Guarda como `assembly/main.go`

## How to Compile

`$ GOARCH=wasm GOOS=js go build`

# For more information

[Link Wiki](https://github.com/AndrusGerman/GoAssembly/wiki)


