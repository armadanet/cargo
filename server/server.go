package server

import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "strconv"
  "github.com/armadanet/cargo/filesystem"
  "log"
  "github.com/armadanet/comms"
)

const (
  Port int    = 8081
  Name string = "armada-storage"
)

type CargoServer interface {
  Run()
  Connect(w http.ResponseWriter, r *http.Request)
  ConnectLoop(socket comms.Socket)
}

type cargoserver struct {
  router  *mux.Router
  filesys filesystem.CargoReadWriter
}

func NewServer() CargoServer {
  return NewCustomServer(filesystem.Hierarchical())
}

func NewCustomServer(filesys filesystem.CargoReadWriter) CargoServer {
  server := &cargoserver{
    router: mux.NewRouter().StrictSlash(true),
    filesys: filesys,
  }
  server.router.HandleFunc("/connect", server.Connect)
  return server
}

func (s *cargoserver) Run() {
  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(Port), s.router))
}

func ConnectSocketAddr() string {
  return fmt.Sprintf("ws://%s:%d/connect", Name, Port)
}
