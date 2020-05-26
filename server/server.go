package server

import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "strconv"
)

type CargoServer interface {
  Run(port int)
  Connect(w http.ResponseWriter, r *http.Request)
}

type cargoserver struct {
  router  *mux.Router
}

func Server() CargoServer {
  server := &cargoserver{
    router: mux.NewRouter().StrictSlash(true),
  }
  server.router.HandleFunc("/connect", server.Connect)
  return server
}



func (s *cargoserver) Run(port int) {
  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), s.router))
}
