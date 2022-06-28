package api

import (
    "github.com/gorilla/mux"
    "simple_go_api/src/domain/usecases"
    "simple_go_api/src/infra/repositories"
    "simple_go_api/src/infra/services"
)

type Config struct {
    Router *mux.Router
    Db     services.DbService
}

func Startup() Config {
    router := mux.NewRouter()

    var dbService = services.NewDbService()
    var bookRepository = repositories.NewBookRepository(dbService)
    var bookUsecase = usecases.NewBookUsecase(bookRepository)

    var ctrls = []Controller{
        NewBookController(bookUsecase),
    }

    for _, ctrl := range ctrls {
        ctrl.Startup(router)
    }

    return Config{
        Router: router,
        Db:     dbService,
    }
}
