package cargo

// https://golang.org/pkg/io/ioutil/
// https://golang.org/pkg/io/

import (
  // "github.com/armadanet/captain/dockercntrl"
  // "io"
  "log"
  "io/ioutil"
)

type Cargo struct {}

func (c *Cargo) Connect() error {
  log.Println("Connect")
  s := New()
  log.Println("About to run")
  s.Run(8081)
  return nil
  // state, _ := dockercntrl.New()
  // err := state.VolumeCreateIdempotent("cargo")
  // return err
}

func (c *Cargo) Get(filename string) ([]byte, error) {
  return ioutil.ReadFile(filename)
}


func (c *Cargo) Put(filename string, data []byte) error {
  return ioutil.WriteFile(filename, data)
}
