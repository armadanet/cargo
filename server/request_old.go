package server

import (
  "github.com/armadanet/comms"
  "log"
)

type Request struct {
  ReqType int
  Name    string
  Data    []byte
}

type Response struct {
  Status  int
  Data    []byte
}

type Requester struct {
  socket *comms.Socket
}

func NewRequester() *Requester {
  socket, err := comms.EstablishSocket("ws://armada-storage:8081/connect")
  if err != nil {
    log.Println(err)
    return nil
  }
  socket.Start(Response{})
  return &Requester{
    socket: &socket,
  }
}

func (r *Requester) SendRequest(req *Request) *Response {
  (*r.socket).Writer() <- *req
  resp := <- (*r.socket).Reader()
  result, ok := resp.(*Response)
  if !ok {
    log.Println("Response wrong type")
    return nil
  }
  return result
}
