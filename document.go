package GoAssembly

import (
	"syscall/js"
)

type documentC struct {
	page  *Page
	Parse parse
}

func Document(t *Page) documentC {
	b := documentC{
		page: t,
	}
	return b
}

func (ap documentC) Alert(tx string) {
	js.Global().Call("alert", tx)
}
