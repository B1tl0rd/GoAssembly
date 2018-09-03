package GoAssembly

import (
	"syscall/js"
)

type parse struct {
}

func (t parse) GetValue(v []js.Value) string {
	IdSelect = v[0].Get("target").Get("id").String()
	return v[0].Get("target").Get("value").String()
}

func (t parse) GetId(v []js.Value) string {
	IdSelect = v[0].Get("target").Get("value").String()
	return v[0].Get("target").Get("id").String()
}
