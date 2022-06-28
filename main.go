package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    "simple_go_api/src/api"
    "syscall"
    "time"
)

func main() {
    var config = api.Startup()

    srv := &http.Server{
        Handler:      config.Router,
        Addr:         "127.0.0.1:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        log.Fatal(srv.ListenAndServe())
    }()

    <-done
    err := config.Db.Dispose()
    if err != nil {
        panic(err)
    }
}
