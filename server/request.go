package server

import (
  "github.com/google/uuid"
  "github.com/armadanet/comms"
)

type ReqType  int
type RespType int

const (
  ReadRequest   ReqType = 1
  WriteRequest  ReqType = 2
)

const (
  SuccessResponse     RespType = 0
  FailureResponse     RespType = -1
  NoSuchFileResponse  RespType = -2
)

type Request struct {
  Id        uuid.UUID
  ReqType   ReqType
  Filename  string
  Data      []byte
}

type NoFilenameError struct {}
func (e *NoFilenameError) Error() string {return "No filename given"}
type NilDataError struct {} // Data can be empty, but not nil
func (e *NilDataError) Error() string {return "Nil data given"}

func NewReadRequest(filename string) (*Request, error) {
  if filename == "" {return nil, &NoFilenameError{}}
  id, err := uuid.NewRandom()
  if err != nil {return nil, err}
  return &Request{
    Id: id,
    ReqType: ReadRequest,
    Filename: filename,
    Data: nil,
  }, nil
}

func NewWriteRequest(filename string, data []byte) (*Request, error) {
  if filename == "" {return nil, &NoFilenameError{}}
  if data == nil {return nil, &NilDataError{}}
  id, err := uuid.NewRandom()
  if err != nil {return nil, err}
  return &Request{
    Id: id,
    ReqType: WriteRequest,
    Filename: filename,
    Data: data,
  }, nil
}

type Response struct {
  Id        uuid.UUID
  RespType  RespType
  Data      []byte
}

type Requestor struct {
  socket          comms.Socket
  RequestChannel  chan *Request
  ResponseChannel chan *Response
}

func NewRequestor() (*Requestor, error) {
  socket, err := comms.EstablishSocket(ConnectSocketAddr())
  if err != nil {return nil, err}
  return CustomRequestor(socket), nil
}

func CustomRequestor(socket comms.Socket) *Requestor {
  r := &Requestor{
    socket: socket,
    RequestChannel: make(chan *Request),
    ResponseChannel: make(chan *Response),
  }
  go func() {
    defer func() {
      close(r.RequestChannel)
      close(r.ResponseChannel)
      r.socket.Close()
    }()
    reader := r.socket.Reader()
    writer := r.socket.Writer()
    for {
      select {
      case input, ok := <- reader:
        if !ok {return}
        data, ok := input.(*Response)
        if !ok {return}
        r.ResponseChannel <- data
      case input, ok := <- r.RequestChannel:
        if !ok {return}
        writer <- input
      }
    }
  }()
  r.socket.Start(Response{})
  return r
}
