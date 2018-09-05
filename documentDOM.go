package GoAssembly

import (
	"os"
	"syscall/js"
)

// errorDOM Contiene una respuesta por error
func errorDOM(errorT string) {
	basiA := `<center> <br/><br/><br/><br/> <h1 style="color:red" >`
	basiB := `<h1/><button onclick="goroute()">Retornar</button></center>`
	GetElementId("app").Set("innerHTML", basiA+errorT+basiB)
	js.Global().Set("ruta", "index")
	os.Exit(1)
}
