package GoAssembly

import (
	"syscall/js"
)

type parse struct {
}

// GetValue del evento se puede odtener el valor del elemento de este evento
func (t parse) GetValue(v []js.Value) string {
	IdSelect = v[0].Get("target").Get("id").String()
	return v[0].Get("target").Get("value").String()
}

// GetId del evento se puede odtener el id
func (t parse) GetId(v []js.Value) string {
	IdSelect = v[0].Get("target").Get("value").String()
	return v[0].Get("target").Get("id").String()
}
