package GoAssembly

//Peticiones Axios
type axios struct {
	page *Page
}

// NewAxios crea un instancia axios
func NewAxios(a *Page) axios {
	return axios{page: a}
}

// Get reques in Axios Js
func (t *axios) Get(url string) string {

	return ""
}
