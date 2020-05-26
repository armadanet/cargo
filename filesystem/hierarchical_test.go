// +build unit

package filesystem_test

import (
  "testing"
  "reflect"
  "github.com/armadanet/cargo/filesystem"
)

func MockFileRead(filename string) ([]byte, error) {
  return []byte(filename), nil
}

func MockFileWrite(filename string, data []byte) error {
  return nil
}

type MockError struct {
  s string
}

func (e *MockError) Error() string {
  return e.s
}

func TestHierarchicalRead(t *testing.T) {
  t.Parallel()
  h := filesystem.CustomHierarchical(MockFileRead, MockFileWrite)
  tests := []struct{
    filename  string
    data      []byte
    err       error
  }{
    {filename: "test.txt", data: []byte("test.txt"), err: nil,},
    {filename: "Hello World", data: []byte("Hello World"), err: nil,},
    {filename: "", data: []byte(""), err: nil,},
  }

  for i, test := range tests {
    data, err := h.ReadFile(test.filename)
    if !reflect.DeepEqual(test.data, data) || !reflect.DeepEqual(err,test.err) {
      t.Errorf("test %d: expected (%v, %v), got (%v, %v)", i+1, test.data, test.err, data, err)
    }
  }
  h = filesystem.CustomHierarchical(func(filename string) ([]byte, error) {
    return []byte(""), &MockError{s: filename}
  }, MockFileWrite)
  data, err := h.ReadFile("my error")
  if (!reflect.DeepEqual([]byte(""), data) || !reflect.DeepEqual(err, &MockError{s: "my error",})) {
    t.Errorf("test: expected (%v, %v), got (%v, %v)", []byte(""), &MockError{s: "my error"}, data, err)
  }
}

func TestHierarchicalWrite(t *testing.T) {
  t.Parallel()
  h := filesystem.CustomHierarchical(MockFileRead, MockFileWrite)
  err := h.WriteFile("test", []byte("test"))
  if err != nil {
    t.Errorf("test: expected (%v), got (%v)", nil, err)
  }
  h = filesystem.CustomHierarchical(MockFileRead, func(filename string, data []byte) error {
    return &MockError{s: "test error",}
  })
  err = h.WriteFile("test", []byte("test"))
  if !reflect.DeepEqual(err, &MockError{s: "test error",}) {
    t.Errorf("test: expected (%v), got (%v)", &MockError{s: "test error",}, err)
  }
}

func TestHierarchicalDefault(t *testing.T) {
  t.Parallel()
  _ = filesystem.Hierarchical()
}
