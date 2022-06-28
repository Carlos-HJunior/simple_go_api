package api

import (
    "net/http"
    "net/http/httptest"
    "simple_go_api/src/infra/services"
)

func NewTestHandler() TestHandler {
    return TestHandler{
        Startup(),
    }
}

type TestHandler struct {
    api Config
}

func (t TestHandler) Dispose() error {
    return t.api.Db.Dispose()
}

func (t TestHandler) Db() services.DbService {
    return t.api.Db
}

func (t *TestHandler) TestRequest(req *http.Request, expectedStatus int) (bool, *httptest.ResponseRecorder) {
    rr := httptest.NewRecorder()
    t.api.Router.ServeHTTP(rr, req)

    return rr.Code == expectedStatus, rr
}
