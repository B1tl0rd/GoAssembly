package GoAssembly

import (
	"os"
	"syscall/js"
)

// errorDOM Contiene una respuesta por error
func errorDOM(errorT string) {
	basiA := `<center> <br/><br/><br/><br/> <h1  class="error">`
	basiB := `<h1/><button class="btn" onclick="goroute()">Retornar</button></center>`

	class := `
	<style>
	.error {color: #D8000C;background-color: #FFD2D2;}
	.btn {
		background-color: #4CAF50; border: none;color: white;
	    padding: 15px 32px;text-align: center;text-decoration: none;
	    display: inline-block;font-size: 16px;}
	</style>
	`
	document.Set("title","Error page")
	GetElementId("app").Set("innerHTML", basiA+errorT+basiB+class)
	js.Global().Set("ruta", "index")
	os.Exit(1)
}
