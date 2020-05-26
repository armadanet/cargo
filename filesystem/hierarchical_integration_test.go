// +build integration

package filesystem_test

import (
  "testing"
  "github.com/armadanet/cargo/filesystem"
  "io/ioutil"
  "os"
  "path/filepath"
)

func TestHierarchicalReadWriteUpdate(t *testing.T) {
  t.Parallel()
  h := filesystem.Hierarchical()
  dir, err := ioutil.TempDir("", "hierarchical_test")
  if err != nil {
    t.Fatalf("ERROR: FAILED TO CREATE TEMP DIR")
  }
  defer os.RemoveAll(dir)
  tmpf := filepath.Join(dir, "tmpfile")

  value1 := "Hello World.\nThis is some basic ascii text.\n"
  if err = h.WriteFile(tmpf, []byte(value1)); err != nil {
    t.Fatalf("Write failed with error: %v", err)
  }

  data, err := h.ReadFile(tmpf)
  if err != nil {
    t.Errorf("Read failed with error: %v", err)
  }
  if string(data) != value1 {
    t.Errorf("Read Mismatch: Expected (%v)\nGot (%v)", value1, string(data))
  }

  value2 := "Updated."
  if err = h.WriteFile(tmpf, []byte(value2)); err != nil {
    t.Fatalf("Write failed with error: %v", err)
  }
  data, err = h.ReadFile(tmpf)
  if err != nil {
    t.Errorf("Read failed with error: %v", err)
  }
  if string(data) != value2 {
    t.Errorf("Read Mismatch: Expected (%v)\nGot (%v)", value2, string(data))
  }
}
