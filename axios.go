package GoAssembly

//Peticiones Axios
type axios struct {
	page *Page
}

// NewAxios crea un instancia axios
func NewAxios(a *Page) axios {
	return axios{page: a}
}

// Get reques Http
func (t *axios) Get(url string) string {

	return ""
}
