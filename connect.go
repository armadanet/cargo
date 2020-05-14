package cargo

import (
  "github.com/armadanet/comms"
  "net/http"
  "log"
  "path/filepath"
  "reflect"
)

// On request adds client through the messenger
func connect() func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    socket, err := comms.AcceptSocket(w,r)
    if err != nil {
      log.Println(err)
      return
    }
    log.Println("Connected to node")
    socket.Start(Request{})
    log.Println("Socket started")
    for {
      select {
      case data, ok := <- socket.Reader():
        if !ok {
          log.Println("Data read corrupted")
          return
        }
        resp, ok := data.(*Request)
        if !ok {
          log.Println("Data read not as a Request")
          log.Println(data)
          log.Println(reflect.TypeOf(data))
          return
        }
        requestedFile := filepath.Join("/data", resp.Name)
        if resp.ReqType == 1 {
          fdata, err := get(requestedFile)
          if err != nil {
            log.Println(err)
            return
          }
          socket.Writer() <- Response{
            Status: 1,
            Data: fdata,
          }
        } else {
          err = put(requestedFile, resp.Data)
          if err != nil {
            log.Println(err)
            return
          }
          socket.Writer() <- Response{
            Status: 0,
            Data: resp.Data,
          }
        }
      }
    }
  }
}
