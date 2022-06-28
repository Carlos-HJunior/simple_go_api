package rules

import (
    "errors"
    "math"
    "simple_go_api/src/domain/dtos"
    "strings"
)

func CreateBookValidation(book dtos.Book) error {
    if len(strings.TrimSpace(book.Name)) == 0 {
        return errors.New("name must not be empty")
    }

    if len(strings.TrimSpace(book.Description)) == 0 {
        return errors.New("description must not be empty")
    }

    if book.Price == 0 || math.Signbit(book.Price) {
        return errors.New("price must not be zero or negative")
    }

    return nil
}
