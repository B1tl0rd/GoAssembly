package GoAssembly

import (
	"syscall/js"
)

//Parse get values from events
type parse struct {
}

// GetValue del evento se puede odtener el valor del elemento de este evento
func (t parse) GetValue(v []js.Value) string {
	IdSelect = v[0].Get("target").Get("id").String()
	//Get
	if GetElementId("null") != v[0].Get("target").Get("value") {
		return v[0].Get("target").Get("value").String()
	}
	return ""
}

// GetId del evento se puede odtener el id
func (t parse) GetId(v []js.Value) string {
	IdSelect = v[0].Get("target").Get("id").String()
	//Get
	if GetElementId("null") != v[0].Get("target").Get("id") {
		return v[0].Get("target").Get("id").String()
	}
	return ""
}
