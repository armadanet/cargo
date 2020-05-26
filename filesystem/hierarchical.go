package filesystem

import (
  "io/ioutil"
)

type ReadFileFunc func(filename string) ([]byte, error)
type WriteFileFunc func(filename string, data []byte) error

type hierarchical struct {
  readfile ReadFileFunc
  writefile WriteFileFunc
}

func Hierarchical() CargoReadWriter {
  return CustomHierarchical(
    ioutil.ReadFile,
    func(filename string, data []byte) error {
      return ioutil.WriteFile(filename, data, 0644)
    })
}

func CustomHierarchical(readfunc ReadFileFunc, writefunc WriteFileFunc) CargoReadWriter {
  return &hierarchical{
    readfile: readfunc,
    writefile: writefunc,
  }
}

func (h *hierarchical) ReadFile(filename string) ([]byte, error) {
  return h.readfile(filename)
}

func (h *hierarchical) WriteFile(filename string, data []byte) error {
  return h.writefile(filename, data)
}
