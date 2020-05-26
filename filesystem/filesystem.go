package filesystem

type CargoReadWriter interface {
  CargoReader
  CargoWriter
}

type CargoReader interface {
  ReadFile(filename string) ([]byte, error)
}

type CargoWriter interface {
  WriteFile(filename string, data []byte) error
}
