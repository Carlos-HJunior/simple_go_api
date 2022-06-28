package usecases

import (
    "context"
    "simple_go_api/src/domain/dtos"
    "simple_go_api/src/domain/rules"
    "simple_go_api/src/ent"
    "simple_go_api/src/infra/repositories"
)

func NewBookUsecase(repository repositories.BookRepository) BookUsecase {
    return bookUsecase{repository}
}

type bookUsecase struct {
    repository repositories.BookRepository
}

func (b bookUsecase) Create(dto dtos.Book, ctx context.Context) (int, error) {
    err := rules.CreateBookValidation(dto)
    if err != nil {
        return 0, err
    }

    return b.repository.CreateOne(ent.Book{
        Name:        dto.Name,
        Description: dto.Description,
        Price:       dto.Price,
    }, ctx)
}

func (b bookUsecase) GetAll(ctx context.Context) ([]dtos.Book, error) {
    items, err := b.repository.GetAll(ctx)
    if err != nil {
        return nil, err
    }

    list := make([]dtos.Book, 0)
    for _, item := range items {
        list = append(list, dtos.Book{
            Name:        item.Name,
            Description: item.Description,
            Price:       item.Price,
        })
    }

    return list, nil
}
