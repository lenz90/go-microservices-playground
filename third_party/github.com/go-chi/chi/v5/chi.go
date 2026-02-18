package chi

import "net/http"

type Router interface {
	http.Handler
	Get(pattern string, h http.HandlerFunc)
	Post(pattern string, h http.HandlerFunc)
	Route(pattern string, fn func(r Router))
	Use(mw ...func(http.Handler) http.Handler)
}

type Mux struct {
	mux *http.ServeMux
	mws []func(http.Handler) http.Handler
}

func NewRouter() *Mux                                           { return &Mux{mux: http.NewServeMux()} }
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) { m.wrap(m.mux).ServeHTTP(w, r) }
func (m *Mux) wrap(h http.Handler) http.Handler {
	for i := len(m.mws) - 1; i >= 0; i-- {
		h = m.mws[i](h)
	}
	return h
}
func (m *Mux) Use(mw ...func(http.Handler) http.Handler) { m.mws = append(m.mws, mw...) }
func (m *Mux) handle(pattern string, h http.HandlerFunc) {
	m.mux.Handle(pattern, m.wrap(http.HandlerFunc(h)))
}
func (m *Mux) Get(pattern string, h http.HandlerFunc) {
	m.handle(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method", 405)
			return
		}
		h(w, r)
	})
}
func (m *Mux) Post(pattern string, h http.HandlerFunc) {
	m.handle(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method", 405)
			return
		}
		h(w, r)
	})
}
func (m *Mux) Route(pattern string, fn func(r Router)) {
	sub := NewRouter()
	fn(sub)
	m.mux.Handle(pattern+"/", http.StripPrefix(pattern, sub))
}
func URLParam(r *http.Request, key string) string { return "" }
