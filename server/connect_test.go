// +build unit

package server_test

import (
  "testing"
  "github.com/armadanet/cargo/server"
  "time"
)

type MockCargoReadWriter struct {
  NextError   error
}

func (rw *MockCargoReadWriter) ReadFile(filename string) ([]byte, error) {
  if rw.NextError != nil {
    err := rw.NextError
    rw.NextError = nil
    return nil, err
  }
  return nil, nil
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
  fsmock := &MockCargoReadWriter{
    NextError: nil,
  }
  s := server.NewCustomServer(fsmock)
  socket := NewMockSocket(t, server.Request{})

  go s.ConnectLoop(socket)
  time.Sleep(2*time.Millisecond)
  if !socket.Started {
    t.Errorf("Failed to start socket")
  }

  

}
