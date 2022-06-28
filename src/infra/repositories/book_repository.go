package repositories

import (
    "context"
    "simple_go_api/src/ent"
    "simple_go_api/src/infra/services"
)

func NewBookRepository(dbService services.DbService) BookRepository {
    return bookRepository{dbService}
}

type bookRepository struct {
    db services.DbService
}

func (r bookRepository) GetAll(ctx context.Context) (ent.Books, error) {
    return r.db.Book.Query().All(ctx)
}

func (r bookRepository) CreateOne(book ent.Book, ctx context.Context) (int, error) {
    data, err := r.db.Book.
        Create().
        SetName(book.Name).
        SetDescription(book.Description).
        SetPrice(book.Price).
        Save(ctx)

    if err != nil {
        return 0, err
    }

    return data.ID, nil
}
