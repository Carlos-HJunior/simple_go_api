package api

import "github.com/gorilla/mux"

type Controller interface {
    Startup(*mux.Router)
}
