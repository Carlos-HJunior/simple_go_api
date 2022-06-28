package repositories

import (
    "context"
    "simple_go_api/src/ent"
)

type BookRepository interface {
    CreateOne(book ent.Book, ctx context.Context) (int, error)
    GetAll(ctx context.Context) (ent.Books, error)
}
