package GoAssembly

import (
	"syscall/js"
)

// documentC cualquier valor del dom
type documentC struct {
	page  *Page
	Parse parse
}

// IdValues valores por defecto de un Id
type IdValues struct {
	Value js.Value
}

// Id contiene los valores del Id select
func (documentC) Id(id string) IdValues {
	el := GetElementId(id)
	if el.String() == "null" {
		errorDOM("El ID:" + id + " No esta disponible")
	}
	return IdValues{
		Value: el.Get("value"),
	}
}

// Document retorna valores especiales referentes al documento
func Document(t *Page) documentC {
	b := documentC{
		page: t,
	}
	return b
}

// Alert Imprime una alerta
func (ap documentC) Alert(tx string) {
	js.Global().Call("alert", tx)
}
