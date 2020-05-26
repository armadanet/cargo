// +build unit

package server_test

import (
  "testing"
  "github.com/armadanet/cargo/server"
  "net/http"
  "net/http/httptest"
)

func TestConnectHandler(t *testing.T) {
  t.Parallel()
  s := server.Server()

  request, err := http.NewRequest(http.MethodGet, "/connect", nil)
  if err != nil {t.Fatalf("Get Create Request Error %v", err)}
  response := httptest.NewRecorder()

  s.Connect(response, request)

  expected := "test"
  result := response.Body.String()

  if expected != result {
    t.Errorf("Expected (%v), got (%v)", expected, result)
  }


}
