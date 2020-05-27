// +build unit

package server_test

import (
  "testing"
  "github.com/armadanet/cargo/server"
)

func TestConnect(t *testing.T) {
  t.Parallel()
  _ = server.NewServer()



}
