package GoAssembly

import (
	"fmt"
	"strconv"
	"syscall/js"
)

type store struct {
	variables map[string]string
}

// Set a la variable tipo 0-> GoAssembli le agrega el valor especificado
func (this *store) Set(name string, val interface{}) {
	name += "Val"
	this.variables[name] = fmt.Sprint(val)
	if Running == true {
		onInnerRunning()
	}

}

// Get odtiene cualquier valor de la variable especificada
func (this *store) Get(name string) string {
	name += "Val"
	return this.variables[name]
}

// GetInt odtiene cualquier valor de la variable especificada en INT
func (this *store) GetInt(name string) int {
	name += "Val"
	n, err := strconv.Atoi(this.variables[name])
	if err != nil {
		return 0
	}
	return n
}

// SetT Guarda las variables GoAssembly tipo 1:false y 2:true
func (this *store) SetT(name string, val interface{}, lvl bool) {
	name += "Val"
	switch lvl {
	case false:
		v := js.Global().Get(name)
		if v == js.Global().Get("undefined") {
			//fmt.Println("No esta definido")
			this.variables[name] = fmt.Sprint(val)
			js.Global().Set(name, fmt.Sprint(val))

		} else {
			//fmt.Println("Si esta definido")
			if Running == false {
				//fmt.Println("Se guarda")
				this.variables[name] = v.String()
			} else {
				//fmt.Println("Se renueva")
				this.variables[name] = fmt.Sprint(val)
				js.Global().Set(name, fmt.Sprint(val))
			}
		}
	case true:
		v := js.Global().Get("localStorage").Call("getItem", name)
		if v == js.Global().Get("localStorage").Call("getItem", "null") {
			//fmt.Println("No esta definido")
			this.variables[name] = fmt.Sprint(val)
			js.Global().Get("localStorage").Call("setItem", name, fmt.Sprint(val))
		} else {
			//fmt.Println("Si esta definido")
			if Running == false {
				//fmt.Println("Se guarda")
				this.variables[name] = v.String()
			} else {
				//fmt.Println("Se renueva")
				this.variables[name] = fmt.Sprint(val)
				js.Global().Get("localStorage").Call("setItem", name, fmt.Sprint(val))
			}
		}
	}
	if Running == true {
		onInnerRunning()
	}
}
