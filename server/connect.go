package server

import (
  "net/http"
  "github.com/armadanet/comms"
  "log"
)

func (s *cargoserver) Connect(w http.ResponseWriter, r *http.Request) {
  socket, err := comms.AcceptSocket(w,r)
  if err != nil {
    log.Println(err)
    return
  }
  s.ConnectLoop(socket)
}

func (s *cargoserver) ConnectLoop(socket comms.Socket) {
  socket.Start(Request{})
  defer func() {
    socket.Close()
  }()
  reader := socket.Reader()
  writer := socket.Writer()
  for {
    input, ok := <- reader
    if !ok {return}
    req, ok := input.(*Request)
    if !ok {return}
    go func() {
      switch req.ReqType {
      case ReadRequest:
        data, err := s.filesys.ReadFile(req.Filename)
        if err != nil {
          writer <- &Response{
            Id: req.Id,
            RespType: FailureResponse,
            Data: []byte(err.Error()),
          }
        } else {
          writer <- &Response{
            Id: req.Id,
            RespType: SuccessResponse,
            Data: data,
          }
        }

      case WriteRequest:
        err := s.filesys.WriteFile(req.Filename, req.Data)
        if err != nil {
          writer <- &Response{
            Id: req.Id,
            RespType: FailureResponse,
            Data: []byte(err.Error()),
          }
        } else {
          writer <- &Response{
            Id: req.Id,
            RespType: SuccessResponse,
            Data: nil,
          }
        }
      }
    }()
  }

}
