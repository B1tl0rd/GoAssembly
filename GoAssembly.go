package GoAssembly

import (
	"os"
	"strings"
	"syscall/js"
)

var Ruta = make(map[string]*Page)
var document = js.Global().Get("document")

//Elemento Base
var dom = "app"

//Ruta por defecto
var ruta = ""

//El Dom a iniciado
var Running = false

//Id de botones rutas
var btnRouteId = []string{}

//Eventos
var eventos []evento

//Guarda los valores y Id de los elementos del DOM
var oldValue = make(map[string]js.Value)

//Guarda el elemento select
var IdSelect string

// Page contiene la base de una pagina
type Page struct {
	// Var Puedes crea variables
	Var store
	// Template contiene el HTML no compilado
	Template string
	// compileTemplate contiene el HTML a mostrar en el DOM
	compileTemplate string

	//->Funciones GoAssembly puedes evitar errores si las utilizas<-
	// Data en esta funcion puedes declarar tu variables
	Data func()
	// Prepare agrega tu codigo que se ejecutara antes del DOM
	Prepare func()
	// Methods Puedes agrega tus metodos -> funciones de GoAssembly
	Methods func()
	// Method agrega los Method declarados
	Method map[string]func([]js.Value)
	// Script codigo que se ejecutar al finalizar la carga del DOM
	Script func()
}

// Title Guarda el titulo de la Pagina
func (m *Page) Title(name string) {
	document.Set("title", name)
}

// RunApp Inicia la paginas declaradas "index" por defecto la ruta principal se declara en el index.htm
func RunApp() {
	rutaMain := js.Global().Get("ruta")
	if rutaMain != js.Global().Get("undefined") {
		rMain := rutaMain.String()
		for na, _ := range Ruta {

			if na == rMain {
				ruta = rMain
				prepareApp(Ruta[rMain])
			}
		}
		errorDOM("No se encontro la ruta -> " + rMain + " <-")

	} else {
		errorDOM("Error la variable ruta no se encuentra disponible")
	}

}

// reloadClickEventRoute recarga los eventos de los botones rutas
func reloadClickEventRoute() {
	for _, v := range btnRouteId {
		GetElementId(v).Call("addEventListener", "click", js.Global().Get("RunRoute"))
	}
}

// prepareApp prepara las funciones webAssembly
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

//processDOM .Algo ?
func processDOM(this *Page) {
	processVARS(this)
	processEVENTS(this)
	Running = true
	processROUTE(this)
}

// processEVENTS  agrega los eventos
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

// processVARS Guarda en el DOM las variables
func processVARS(this *Page) {
	tmp := this.Template
	for name, val := range this.Var.variables {
		var a = len(name)
		nomA := "{{ " + name[:a-3] + " }}"
		nomB := "{{" + name + "}}"
		nomC := "{{ " + name + "}}"
		nomD := "{{" + name + " }}"
		rem := strings.NewReplacer(nomA, val, nomB, val, nomC, val, nomD, val)
		tmp = rem.Replace(tmp)
	}
	this.compileTemplate = tmp
}

// RuntRoute detiene la app y la inicia en la ruta declarada
func RunRoute(laruta string) {
	js.Global().Set("ruta", laruta)
	js.Global().Call("goroute")
	os.Exit(0)
}

//onInnerRunning si la pagina se renueva recarga los elementos
func onInnerRunning() {
	getIdTotalValue()
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
	if GetElementId(IdSelect) != GetElementId("null") {
		GetElementId(IdSelect).Call("focus")
	}
}

// processROUTEs Procesa las rutas
func processROUTE(this *Page) {
	var c = make(chan struct{}, 0)
	//Crea un canal de eventos
	idbtn := setIdRoute(this)
	HTMLset(dom, this.compileTemplate)
	//Agrega

	funct := func(t []js.Value) {
		p := t[0].Get("target").Get("id").String()
		RunRoute(getRuta(p))
	}
	js.Global().Set("RunRoute", js.NewCallback(funct))
	//Events click
	for _, v := range idbtn {
		btnRouteId = append(btnRouteId, v)
		GetElementId(v).Call("addEventListener", "click", js.Global().Get("RunRoute"))
	}
	//agrega los eventos
	reloadEvents()
	//Finaliza
	if this.Script != nil {
		go this.Script()
	}
	<-c
}

// reloadEvents recarga los eventos
func reloadEvents() {
	for _, n := range eventos {
		if GetElementId(n.MiId) != GetElementId("null") {
			GetElementId(n.MiId).Call("addEventListener", n.evento, js.Global().Get("Event"+n.methodo))
		} else {
			//No definido
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

// HTMLset al id especificado iserta
func HTMLset(name, html string) {
	GetElementId(name).Set("innerHTML", html)
}

// GetElementId llama el dato mediante el Id
func GetElementId(name string) js.Value {
	return js.Global().Get("document").Call("getElementById", name)
}
