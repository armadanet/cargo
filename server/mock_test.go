// +build unit integration

package server_test

import (
  "testing"
  "reflect"
  "github.com/armadanet/cargo/server"
  "time"
)

type MockSocket struct {
  Read        chan interface{}
  Write       chan interface{}
  ReqLog      []server.Request
  RespLog     []server.Response
  startType   reflect.Type
  t           *testing.T
  Started     bool
  CloseCalled bool
  quit        chan interface{}
}

func NewMockSocket(t *testing.T, read interface{}) *MockSocket {
  return &MockSocket{
    t: t,
    Write: make(chan interface{}),
    Read: make(chan interface{}),
    ReqLog: []server.Request{},
    RespLog: []server.Response{},
    startType: reflect.TypeOf(read),
    Started: false,
    CloseCalled: false,
    quit: make(chan interface{}),
  }
}
func (s *MockSocket) Reader() chan interface{} {return s.Read}
func (s *MockSocket) Writer() chan interface{} {return s.Write}
func (s *MockSocket) Start(read interface{}) {
  if reflect.TypeOf(read) != s.startType {
    s.t.Errorf("Started with (%T), not (%v)", read, s.startType)
  }
  go func() {
    for {
      select {
      case input, ok := <- s.Write:
        if !ok {
          s.t.Errorf("Write Closed")
        }
        switch v := input.(type) {
        case *server.Request:
          s.ReqLog = append(s.ReqLog, *v)
        case *server.Response:
          s.RespLog = append(s.RespLog, *v)
        default:
          s.t.Errorf("Malformed request (%v)", input)
        }
      case <- s.quit:
        return
      }
    }
  }()
  s.Started = true
}
func (s *MockSocket) Close() {
  s.CloseCalled = true
  if !s.CloseCalled {
    close(s.quit)
    time.Sleep(2*time.Millisecond)
    close(s.Write)
    close(s.Read)
  }
}

type MockCargoReadWriter struct {
  NextError     error
  FileRead      string
  DataRecieved  []byte
  DataSend      []byte
}

func NewMockCargoReadWriter() *MockCargoReadWriter {
  return &MockCargoReadWriter{
    NextError: nil,
    FileRead: "",
    DataRecieved: nil,
    DataSend: nil,
  }
}

func (rw *MockCargoReadWriter) ReadFile(filename string) ([]byte, error) {
  if rw.NextError != nil {
    err := rw.NextError
    rw.NextError = nil
    return nil, err
  }
  return []byte(filename), nil
}

func (rw *MockCargoReadWriter) WriteFile(filename string, data []byte) error {
  if rw.NextError != nil {
    err := rw.NextError
    rw.NextError = nil
    return err
  }
  return nil
}
