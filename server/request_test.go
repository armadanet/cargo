// +build unit

package server_test

import (
  "testing"
  "github.com/armadanet/cargo/server"
  "github.com/google/uuid"
  "reflect"
  "time"
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
  socket := NewMockSocket(t, server.Response{})
  requestor := server.CustomRequestor(socket)
  if !socket.Started {
    t.Errorf("Failed to start socket")
  }
  select {
  case request := <- requestor.RequestChannel:
    t.Errorf("Channel should be empty, got (%v)", request)
  case response := <- requestor.ResponseChannel:
    t.Errorf("Channel should be empty, got (%v)", response)
  case <- time.After(2*time.Millisecond):
  }

  req1, _ := server.NewReadRequest("test1")
  req2, _ := server.NewWriteRequest("test2", []byte("test2"))
  requestor.RequestChannel <- req1
  select {
  case request := <- requestor.RequestChannel:
    t.Errorf("Channel should be empty, got (%v)", request)
  case response := <- requestor.ResponseChannel:
    t.Errorf("Channel should be empty, got (%v)", response)
  case <- time.After(2*time.Millisecond):
  }
  requestor.RequestChannel <- req2

  resp := &server.Response{
    Id: req1.Id,
    RespType: server.SuccessResponse,
    Data: nil,
  }
  socket.Read <- resp
  select {
  case request := <- requestor.RequestChannel:
    t.Errorf("Channel should have response, got (%v)", request)
  case response := <- requestor.ResponseChannel:
    if !reflect.DeepEqual(resp, response) {
      t.Errorf("Expected (%v), got (%v)", resp, response)
    }
  case <- time.After(2*time.Millisecond):
    t.Errorf("Channel should not be empty")
  }
  time.Sleep(2*time.Millisecond)

  if !reflect.DeepEqual(socket.ReqLog, []server.Request{*req1, *req2}) {
    t.Errorf("Expected (%v), got (%v)", []server.Request{*req1, *req2}, socket.ReqLog)
  }

  if !reflect.DeepEqual(socket.RespLog, []server.Response{}) {
    t.Errorf("Expected empty response log, got (%v)", socket.RespLog)
  }

  if socket.CloseCalled {
    t.Errorf("Socket closed prematurely")
  }

  requestor.Quit()
  time.Sleep(2*time.Millisecond)
  if !socket.CloseCalled {
    t.Errorf("Failed to end socket connection")
  }
}

func TestRequestorImproperClosing(t *testing.T) {
  t.Parallel()
  socket := NewMockSocket(t, server.Response{})
  requestor := server.CustomRequestor(socket)
  select {
  case <- requestor.RequestChannel:
    t.Errorf("Request Channel closed prematurely")
  default:
  }
  select {
  case <- requestor.ResponseChannel:
    t.Errorf("Response Channel closed prematurely")
  default:
  }

  close(socket.Read)
  time.Sleep(2*time.Millisecond)
  select {
  case <- requestor.RequestChannel:
  default:
    t.Errorf("Request Channel unclosed")
  }
  select {
  case <- requestor.ResponseChannel:
  default:
    t.Errorf("Response Channel unclosed")
  }

  socket = NewMockSocket(t, server.Response{})
  requestor = server.CustomRequestor(socket)
  select {
  case <- requestor.RequestChannel:
    t.Errorf("Request Channel closed prematurely")
  default:
  }
  select {
  case <- requestor.ResponseChannel:
    t.Errorf("Response Channel closed prematurely")
  default:
  }

  close(requestor.RequestChannel)
  time.Sleep(2*time.Millisecond)
  select {
  case <- requestor.RequestChannel:
  default:
    t.Errorf("Request Channel unclosed")
  }
  select {
  case <- requestor.ResponseChannel:
  default:
    t.Errorf("Response Channel unclosed")
  }
}
