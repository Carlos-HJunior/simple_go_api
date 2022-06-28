package api

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "net/http"
    "simple_go_api/src/domain/dtos"
    "testing"
)

func TestCreateBook(t *testing.T) {
    var handler = NewTestHandler()

    var item = dtos.Book{
        Name:        "jorge",
        Description: "desc",
        Price:       12.54,
    }

    byts, err := json.Marshal(item)
    if err != nil {
        t.Errorf("format error: %v", err)
        return
    }

    request, err := http.NewRequest(http.MethodPost, "/book/create", bytes.NewBuffer(byts))
    if err != nil {
        assert.Error(t, err, "request error")
        return
    }

    ok, _ := handler.TestRequest(request, http.StatusOK)
    if !ok {
        assert.Error(t, err, "request error")
        return
    }

    all, err := handler.Db().Book.Query().All(context.TODO())
    if err != nil {
        assert.Error(t, err, "database query error")
        return
    }

    assert.NotZero(t, len(all), "book not created")
    assert.Equal(t, item, dtos.Book{
        Name:        all[0].Name,
        Description: all[0].Description,
        Price:       all[0].Price,
    }, "book not equal")

    handler.Dispose()
}
