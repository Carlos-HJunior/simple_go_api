package api

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "io"
    "net/http"
    "simple_go_api/src/domain/dtos"
    "simple_go_api/src/domain/usecases"
)

func NewBookController(usecase usecases.BookUsecase) Controller {
    return bookController{usecase}
}

type bookController struct {
    usecase usecases.BookUsecase
}

func (c bookController) Startup(r *mux.Router) {
    prefix := r.PathPrefix("/book").Subrouter()

    prefix.HandleFunc("/create", c.post).Methods(http.MethodPost)
    prefix.HandleFunc("/all", c.getAll).Methods(http.MethodGet)
}

func (c bookController) post(w http.ResponseWriter, r *http.Request) {
    bytes, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "body error", http.StatusBadRequest)
        return
    }

    var dto = dtos.Book{}
    err = json.Unmarshal(bytes, &dto)
    if err != nil {
        http.Error(w, "body error", http.StatusBadRequest)
        return
    }

    id, err := c.usecase.Create(dto, r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, _ = fmt.Fprint(w, id)
}

func (c bookController) getAll(w http.ResponseWriter, r *http.Request) {
    items, err := c.usecase.GetAll(r.Context())
    if err != nil {
        http.Error(w, "", http.StatusInternalServerError)
        return
    }

    bytes, err := json.Marshal(items)
    if err != nil {
        http.Error(w, "", http.StatusInternalServerError)
        return
    }

    _, _ = w.Write(bytes)
}
