package GoAssembly

import (
	"strconv"
	"strings"
	"syscall/js"
)

var Ruta = make(map[string]*Page)
var document = js.Global().Get("document")
var dom = "app"
var ruta = ""
var Running = false
var btnRouteId = []string{}
var eventos []evento
var totalId = []string{}
var oldValue = make(map[string]js.Value)
var IdSelect string

type Page struct {
	Var             store
	Template        string
	compileTemplate string
	Data            func()
	Prepare         func()
	Methods         func()
	Method          map[string]func([]js.Value)
	Script          func()
}

func (m *Page) Title(name string) {
	document.Set("title", name)
}

func RunApp() {
	rutaMain := js.Global().Get("ruta").String()
	if rutaMain != "undefined" {
		for na, _ := range Ruta {
			if na == rutaMain {
				ruta = rutaMain
				prepareApp(Ruta[rutaMain])
			}
		}
		errorDOM("No se encontro la ruta -> " + rutaMain + " <-")

	} else {
		errorDOM("Error la variable ruta no se encuentra disponible")
	}

}

func reloadClickEventRoute() {
	for _, v := range btnRouteId {
		GetElementId(v).Call("addEventListener", "click", js.Global().Get("RunRoute"))
	}
}

func prepareApp(app *Page) {
	app.Var.variables = make(map[string]string)
	app.Method = make(map[string]func([]js.Value))
	if app.Data != nil {
		app.Data()
	}
	if app.Prepare != nil {
		app.Prepare()
	}
	if app.Methods != nil {
		app.Methods()
	}
	processDOM(app)
}

//Dom proccess
func processDOM(this *Page) {
	processVARS(this)
	processEVENTS(this)
	Running = true
	processROUTE(this)
}

func processEVENTS(this *Page) {
	//se agrega la func
	for n, v := range this.Method {
		js.Global().Set("Event"+n, js.NewCallback(v))
	}
	setIdAndEventName(this)
}

type evento struct {
	evento  string
	methodo string
	old     string
	newE    string
	iden    string
	MiId    string
}

//Retorna el id y el nombre del evento
func setIdAndEventName(template *Page) {
	prefi := `@e=`
	lenPre := len(prefi) + 1
	//Guarda los datos
	trab := false
	poxA := 0
	poxB := 0
	Numero := 0

	for i, v := range template.compileTemplate {
		if string(v) == "\"" {
			if trab == true {
				trab = false
				poxB = i
				if template.compileTemplate[poxA-lenPre+1:poxA] == prefi {
					//Old data
					old := template.compileTemplate[poxA : poxB+1]
					oldEl := prefi + old
					agrega := "Ev" + strconv.Itoa(Numero)
					parse := template.compileTemplate[poxA+1 : poxB]
					n := strings.LastIndex(parse, ".")
					//fmt.Println(parse[n+1:])
					//fmt.Println(parse[:n])
					eventos = append(eventos, evento{
						old:     oldEl,
						evento:  parse[n+1:],
						methodo: parse[:n],
						iden:    agrega,
					})

					Numero++
				}
			} else {
				trab = true
				poxA = i
			}
		}
	}
	for i, v := range eventos {
		eventos[i].newE = "id=\"" + v.methodo + v.evento + v.iden + "\""
		eventos[i].MiId = v.methodo + v.evento + v.iden
		v = eventos[i]
		template.compileTemplate = strings.Replace(template.compileTemplate, v.old, v.newE, 1)
	}
	GetElementId("app").Set("innerHTML", template.compileTemplate)
}

//Guarda en el DOM las variables
func processVARS(this *Page) {
	tmp := this.Template
	for name, val := range this.Var.variables {
		var a = len(name)
		nomA := "{{ " + name[:a-3] + " }}"
		nomB := "{{" + name + "}}"
		rem := strings.NewReplacer(nomA, val, nomB, val)
		tmp = rem.Replace(tmp)
	}
	this.compileTemplate = tmp
}

func onInnerRunning() {
	saveOldValue()
	processVARS(Ruta[ruta])
	GetElementId("app").Set("innerHTML", Ruta[ruta].compileTemplate)
	setIdRoute(Ruta[ruta])
	GetElementId("app").Set("innerHTML", Ruta[ruta].compileTemplate)
	for i, v := range eventos {
		eventos[i].newE = "id=\"" + v.methodo + v.evento + v.iden + "\""
		eventos[i].MiId = v.methodo + v.evento + v.iden
		v = eventos[i]
		Ruta[ruta].compileTemplate = strings.Replace(Ruta[ruta].compileTemplate, v.old, v.newE, 1)
	}
	GetElementId("app").Set("innerHTML", Ruta[ruta].compileTemplate)
	getOldValue()
	reloadEvents()
	reloadClickEventRoute()
	if GetElementId(IdSelect).String() != "null" {
		GetElementId(IdSelect).Call("focus")
	}
}

//Procesa las rutas
func processROUTE(this *Page) {
	var c = make(chan struct{}, 0)
	//Crea un canal de eventos
	idbtn := setIdRoute(this)
	HTMLset(dom, this.compileTemplate)
	//Agrega

	funct := func(t []js.Value) {
		p := t[0].Get("target").Get("id").String()
		js.Global().Set("ruta", getRuta(p))
		js.Global().Call("goroute")
		panic("")
	}
	js.Global().Set("RunRoute", js.NewCallback(funct))
	//Events click
	for _, v := range idbtn {
		btnRouteId = append(btnRouteId, v)
		GetElementId(v).Call("addEventListener", "click", js.Global().Get("RunRoute"))
	}
	//agrega los eventos
	reloadEvents()
	//GUarda todo los ID
	for _, v := range idbtn {
		totalId = append(totalId, v)
	}
	for _, v2 := range eventos {
		totalId = append(totalId, v2.MiId)
	}
	//Finaliza
	if this.Script != nil {
		go this.Script()
	}
	<-c
}

func reloadEvents() {
	for _, n := range eventos {
		if GetElementId(n.MiId).String() != "null" {
			GetElementId(n.MiId).Call("addEventListener", n.evento, js.Global().Get("Event"+n.methodo))
		} else {
			println("Error no esta defido")
		}
	}
}

//Se odtiene la ruta mediante el id
func getRuta(idbtn string) string {
	for i, _ := range idbtn {
		if idbtn[i:i+2] == "Rt" {
			return idbtn[:i]
		}
	}
	return "Error"
}

//Convierte los elementos rutas en id
func setIdRoute(template *Page) []string {
	prefi := `@rt=`
	newPrefi := `id=`
	lenPre := len(prefi) + 1
	datos := []string{}
	remplazar := [][]string{}
	trab := false
	poxA := 0
	poxB := 0
	Numero := 0
	for i, v := range template.compileTemplate {
		if string(v) == "\"" {
			if trab == true {
				trab = false
				poxB = i
				if template.compileTemplate[poxA-lenPre+1:poxA] == prefi {
					//Old data
					old := template.compileTemplate[poxA : poxB+1]
					oldEl := prefi + old
					newEl := newPrefi + template.compileTemplate[poxA:poxB] + "Rt" + strconv.Itoa(Numero) + "\""
					remplazar = append(remplazar, []string{oldEl, newEl})
					datos = append(datos, template.compileTemplate[poxA+1:poxB]+"Rt"+strconv.Itoa(Numero))
					Numero++
				}
			} else {
				trab = true
				poxA = i
			}
		}
	}
	for _, v := range remplazar {
		template.compileTemplate = strings.Replace(template.compileTemplate, v[0], v[1], 1)
	}
	return datos
}

func HTMLset(name, html string) {
	GetElementId(name).Set("innerHTML", html)
}

// GetElementId llama el dato mediante el Id
func GetElementId(name string) js.Value {
	return js.Global().Get("document").Call("getElementById", name)
}

func saveOldValue() {
	oldValue = nil
	oldValue = make(map[string]js.Value)
	for _, V := range totalId {
		if GetElementId(V).String() != "null" {
			oldValue[V] = GetElementId(V).Get("value")
		}
	}
}

func getOldValue() {
	for name, valor := range oldValue {
		if GetElementId(name).String() != "null" {
			GetElementId(name).Set("value", valor)
		}
	}
}
