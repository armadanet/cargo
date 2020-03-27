package cargo

// https://golang.org/pkg/io/ioutil/
// https://golang.org/pkg/io/

import (
  // "github.com/armadanet/captain/dockercntrl"
  // "io"
  "fmt"
)

type Cargo struct {}

func (c *Cargo) Connect() error {
  fmt.Println("Connect")
  for {}
  return nil
  // state, _ := dockercntrl.New()
  // err := state.VolumeCreateIdempotent("cargo")
  // return err
}

func (c *Cargo) Get(filename string) []byte {
  return []byte{}
}


func (c *Cargo) Put(filename string, data []byte) {}
