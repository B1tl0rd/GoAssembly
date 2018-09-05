package GoAssembly

import (
	"syscall/js"
)

// documentC cualquier valor del dom
type documentC struct {
	page *Page
	//Parse
	Parse parse
}

// IdValues valores por defecto de un Id
type IdValues struct {
	Value js.Value
}

// Id contiene los valores del Id select
func (documentC) Id(id string) IdValues {
	el := GetElementId(id)
	if el == GetElementId("null") {
		errorDOM("El ID:" + id + " No esta disponible")
	}
	//Get Values from Id
	return IdValues{
		Value: el.Get("value"),
	}
}

// Document retorna valores especiales referentes al documento
func Document(t *Page) documentC {
	//Get document
	b := documentC{
		page: t,
	}
	return b
}

// Alert Imprime una alerta
func (ap documentC) Alert(tx string) {
	js.Global().Call("alert", tx)
}

/*
// SetHTML
func (a *documentC) SetHTML(id string, html string) {
	HTMLset(id, element)
}
*/

// Log console.log js
func (a *documentC) Log(data js.Value) {
	js.Global().Get("console").Call("log", data)
}
