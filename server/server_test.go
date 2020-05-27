// +build unit

package server_test

import (
  "testing"
  "github.com/armadanet/cargo/server"
)

func TestServerSetup(t *testing.T) {
  t.Parallel()
  _ = server.NewServer()

  expected := "ws://armada-storage:8081/connect"
  got := server.ConnectSocketAddr()
  if expected != got {
    t.Errorf("Expected (%v), got (%v)", expected, got)
  }
}
