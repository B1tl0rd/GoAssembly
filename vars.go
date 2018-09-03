package GoAssembly

import (
	"fmt"
	"strconv"
	"syscall/js"
)

type store struct {
	variables map[string]string
}

func (this *store) Set(name string, val interface{}) {
	name += "Val"
	this.variables[name] = fmt.Sprint(val)
	if Running == true {
		onInnerRunning()
	}

}

func (this *store) Get(name string) string {
	name += "Val"
	return this.variables[name]
}
func (this *store) GetInt(name string) int {
	name += "Val"
	n, err := strconv.Atoi(this.variables[name])
	if err != nil {
		return 0
	}
	return n
}

func (this *store) SetT(name string, val interface{}, lvl bool) {
	name += "Val"
	switch lvl {
	case false:
		v := js.Global().Get(name).String()
		if v == "undefined" {
			//fmt.Println("No esta definido")
			this.variables[name] = fmt.Sprint(val)
			js.Global().Set(name, fmt.Sprint(val))

		} else {
			//fmt.Println("Si esta definido")
			if Running == false {
				//fmt.Println("Se guarda")
				this.variables[name] = v
			} else {
				//fmt.Println("Se renueva")
				this.variables[name] = fmt.Sprint(val)
				js.Global().Set(name, fmt.Sprint(val))
			}
		}
	case true:
		v := js.Global().Get("localStorage").Call("getItem", name).String()
		if v == "null" {
			//fmt.Println("No esta definido")
			this.variables[name] = fmt.Sprint(val)
			js.Global().Get("localStorage").Call("setItem", name, fmt.Sprint(val))
		} else {
			//fmt.Println("Si esta definido")
			if Running == false {
				//fmt.Println("Se guarda")
				this.variables[name] = v
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
