package cargo

import (
  "github.com/armadanet/comms"
  "net/http"
  "log"
)

// On request adds client through the messenger
func connect() func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    _, err := comms.AcceptSocket(w,r)
    if err != nil {
      log.Println(err)
      return
    }
    log.Println("Connected")
  }
}
