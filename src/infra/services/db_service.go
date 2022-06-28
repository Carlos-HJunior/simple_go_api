package services

import (
    "context"
    _ "github.com/mattn/go-sqlite3"
    "simple_go_api/src/ent"
)

func NewDbService() DbService {
    client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        panic(err)
    }

    err = client.Schema.Create(context.Background())
    if err != nil {
        panic(err)
    }

    return DbService{client}
}

type DbService struct {
    *ent.Client
}

func (s DbService) Dispose() error {
    return s.Close()
}
