package usecases

import (
    "context"
    "simple_go_api/src/domain/dtos"
)

type BookUsecase interface {
    Create(dtos.Book, context.Context) (int, error)
    GetAll(ctx context.Context) ([]dtos.Book, error)
}
