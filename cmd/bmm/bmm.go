package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"


    "github.com/tomlitton/boggle-my-mind/pkg/api"
)

func httpRouter() (*mux.Router) {
    router := mux.NewRouter()

    router.HandleFunc("/board/{rows}/{columns}/{letters}", api.HandleBoard).Methods("GET")
    router.HandleFunc("/health", api.Health).Methods("GET")

    return router
}

func main() {
    // @TODO: Pass file as cli arg or env variable
    api.LoadData("assets/dictionary.txt")
    log.Printf("Starting http server")
    log.Fatal(http.ListenAndServe(":8080", httpRouter()))
}