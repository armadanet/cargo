package server

import (
  "net/http"
  "github.com/armadanet/comms"
)


func (s *cargoserver) Connect(w http.ResponseWriter, r *http.Request) {
  socket, err := comms.AcceptSocket(w,r)
  if err != nil {
    log.Println(err)
    return
  }
  

}
