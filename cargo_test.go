package cargo_test

import (
  "testing"
  "github.com/armadanet/cargo/filesystem"
)

func MockFileRead(filename string) ([]byte, error) {
  return []byte(filename), nil
}

func MockFileWrite(filename string, data []byte) error {
  return nil
}

func TestCargo(t *testing.T) {
  h := filesystem.CustomHierarchical(MockFileRead, MockFileWrite)
  if h.WriteFile("test.txt", []byte("data\n")) != nil {
    t.Errorf("Mock file write returned error")
  }
}
