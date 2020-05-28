// +build integration

package server_test

import (
  "testing"
  "github.com/armadanet/cargo/server"
  "time"
)

func TestServerConnection(t *testing.T) {
  fsmock := NewMockCargoReadWriter()
  s := server.NewCustomServer(fsmock)

  go s.Run()
  time.Sleep(3*time.Second)
}
