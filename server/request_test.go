// +build unit

package server_test

import (
  "testing"
  "github.com/armadanet/cargo/server"
  "github.com/google/uuid"
  "reflect"
)

func TestNewReadRequestError(t *testing.T) {
  t.Parallel()
  req, err := server.NewReadRequest("")
  _, ok := err.(*server.NoFilenameError); if !ok || req != nil {
    t.Errorf("Expected (%v, %v), got (%v, %v)", nil, &server.NoFilenameError{},
      req, err)
  }
}

func TestNewReadRequestFormat(t *testing.T) {
  t.Parallel()
  filename := "myfile"
  req, err := server.NewReadRequest(filename)
  if req == nil || err != nil {
    t.Errorf("Expected (not nil, nil), got (%v, %v)", req, err)
  }
  if reflect.DeepEqual(uuid.UUID{}, req.Id) {
    t.Errorf("Id unset")
  }
  if uuid.New().Version() != req.Id.Version() {
    t.Errorf("Expected Id (%v), got (%v)", uuid.New().Version().String(),
      req.Id.Version().String())
  }
  if req.ReqType != server.ReadRequest {
    t.Errorf("Expected (%v), got (%v)", server.ReadRequest, req.ReqType)
  }
  if req.Filename != filename {
    t.Errorf("Expected (%v), got (%v)", filename, req.Filename)
  }
  if req.Data != nil {
    t.Errorf("Expected (<nil>), got (%v)", req.Data)
  }
}

func TestNewWriteRequestError(t *testing.T) {
  t.Parallel()
  req, err := server.NewWriteRequest("", []byte("data"))
  _, ok := err.(*server.NoFilenameError); if !ok || req != nil {
    t.Errorf("Expected (%v, %v), got (%v, %v)", nil, &server.NoFilenameError{},
      req, err)
  }

  req, err = server.NewWriteRequest("filename", nil)
  _, ok = err.(*server.NilDataError); if !ok || req != nil {
    t.Errorf("Expected (%v, %v), got (%v, %v)", nil, &server.NoFilenameError{},
      req, err)
  }

  req, err = server.NewWriteRequest("", nil)
  switch err.(type) {
  case *server.NoFilenameError:
  case *server.NilDataError:
  default:
    t.Errorf("Unexpected error (%v)", err)
  }
}

func TestNewWriteRequestFormat(t *testing.T) {
  t.Parallel()
  filename := "myfile"
  data := []byte("my data")

  req, err := server.NewWriteRequest(filename, data)
  if req == nil || err != nil {
    t.Errorf("Expected (not nil, nil), got (%v, %v)", req, err)
  }

  if reflect.DeepEqual(uuid.UUID{}, req.Id) {
    t.Errorf("Id unset")
  }
  if uuid.New().Version() != req.Id.Version() {
    t.Errorf("Expected Id (%v), got (%v)", uuid.New().Version().String(),
      req.Id.Version().String())
  }
  if req.ReqType != server.WriteRequest {
    t.Errorf("Expected (%v), got (%v)", server.WriteRequest, req.ReqType)
  }
  if req.Filename != filename {
    t.Errorf("Expected (%v), got (%v)", filename, req.Filename)
  }
  if !reflect.DeepEqual(req.Data, data) {
    t.Errorf("Expected (%v), got (%v)", data, req.Data)
  }

  req, err = server.NewWriteRequest(filename, []byte(""))
  if req == nil || err != nil {
    t.Errorf("Expected (not nil, nil), got (%v, %v)", req, err)
  }
  if !reflect.DeepEqual(req.Data, []byte("")) {
    t.Errorf("Expected (%v), got (%v)", []byte(""), req.Data)
  }
}

func TestCustomRequestor(t *testing.T) {
  t.Parallel()
  // type MockSocket struct {
  //
  // }
  // func (s *MockSocket) Reader()
}
