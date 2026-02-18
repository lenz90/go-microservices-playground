package main

import (
	"net/http"
	"playground/03-microservice-foundation/01-clean-arch/internal/repository/inmem"
	httptransport "playground/03-microservice-foundation/01-clean-arch/internal/transport/http"
	"playground/03-microservice-foundation/01-clean-arch/internal/usecase"
)

func main() {
	r := inmem.New()
	s := usecase.Service{R: r}
	_ = http.ListenAndServe(":8080", httptransport.Router(s))
}
