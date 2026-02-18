package httpx

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Handler struct{ users, payments *url.URL }

func New(usersBase, paymentsBase string) (*Handler, error) {
	u, err := url.Parse(usersBase)
	if err != nil {
		return nil, err
	}
	p, err := url.Parse(paymentsBase)
	if err != nil {
		return nil, err
	}
	return &Handler{users: u, payments: p}, nil
}

func (h *Handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/users", httputil.NewSingleHostReverseProxy(h.users))
	mux.Handle("/payments", httputil.NewSingleHostReverseProxy(h.payments))
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte("ok")) })
	return mux
}
