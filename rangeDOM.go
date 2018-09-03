package GoAssembly

import (
	"strconv"
	"strings"
)

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

func getOldValue() {
	for name, valor := range oldValue {
		if GetElementId(name).String() != "null" {
			GetElementId(name).Set("value", valor)
		}
	}
}

func getIdTotalValue() {
	template := Ruta[ruta]
	prefi := `id=`
	lenPre := len(prefi) + 1
	//Guarda los datos
	trab := false
	poxA := 0
	poxB := 0
	for i, v := range template.compileTemplate {
		if string(v) == "\"" {
			if trab == true {
				trab = false
				poxB = i
				if template.compileTemplate[poxA-lenPre+1:poxA] == prefi {
					//Old data
					old := template.compileTemplate[poxA+1 : poxB]
					if GetElementId(old).String() != "null" {
						oldValue[old] = GetElementId(old).Get("value")
					}

				}
			} else {
				trab = true
				poxA = i
			}
		}
	}
}
