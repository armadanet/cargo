// +build unit

package server_test

import (
  "testing"
  "github.com/armadanet/cargo/server"
  "time"
  "errors"
  "reflect"
)

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


func TestConnectLoop(t *testing.T) {
  t.Parallel()
  fsmock := NewMockCargoReadWriter()
  s := server.NewCustomServer(fsmock)
  socket := NewMockSocket(t, server.Request{})

  go s.ConnectLoop(socket)
  time.Sleep(2*time.Millisecond)
  if !socket.Started {
    t.Errorf("Failed to start socket")
  }

  req1, _ := server.NewReadRequest("test1")
  socket.Read <- req1
  time.Sleep(2*time.Millisecond)
  req1err := errors.New("test1 error")
  fsmock.NextError = req1err
  socket.Read <- req1
  time.Sleep(2*time.Millisecond)

  req2, _ := server.NewWriteRequest("test2", []byte("test2"))
  socket.Read <- req2
  time.Sleep(2*time.Millisecond)
  req2err := errors.New("test2 error")
  fsmock.NextError = req2err
  socket.Read <- req2
  time.Sleep(2*time.Millisecond)

  responses := []server.Response{
    server.Response{Id: req1.Id, RespType: server.SuccessResponse, Data: []byte("test1"),},
    server.Response{Id: req1.Id, RespType: server.FailureResponse, Data: []byte("test1 error"),},
    server.Response{Id: req2.Id, RespType: server.SuccessResponse, Data: nil,},
    server.Response{Id: req2.Id, RespType: server.FailureResponse, Data: []byte("test2 error"),},
  }

  if !reflect.DeepEqual(socket.ReqLog, []server.Request{}) {
    t.Errorf("Expected empty, got (%v)", socket.ReqLog)
  }

  if !reflect.DeepEqual(socket.RespLog, responses) {
    t.Errorf("Expected (%v), got (%v)", responses, socket.RespLog)
  }
}

func TestConnectLoopFailure(t *testing.T) {
  t.Parallel()
  fsmock := NewMockCargoReadWriter()
  s := server.NewCustomServer(fsmock)
  socket := NewMockSocket(t, server.Request{})

  go s.ConnectLoop(socket)
  time.Sleep(2*time.Millisecond)
  if !socket.Started {
    t.Errorf("Failed to start socket")
  }

  close(socket.Read)
  time.Sleep(2*time.Millisecond)
  if !socket.CloseCalled {
    t.Errorf("Socket not closed")
  }

  socket = NewMockSocket(t, server.Request{})
  go s.ConnectLoop(socket)
  time.Sleep(2*time.Millisecond)
  if !socket.Started {
    t.Errorf("Failed to start socket")
  }

  socket.Read <- "hello"
  time.Sleep(2*time.Millisecond)
  if !socket.CloseCalled {
    t.Errorf("Socket not closed")
  }

}
