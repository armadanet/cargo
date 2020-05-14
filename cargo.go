package cargo

// https://golang.org/pkg/io/ioutil/
// https://golang.org/pkg/io/

import (
  // "github.com/armadanet/captain/dockercntrl"
  // "io"
  "log"
)

type Cargo struct {}

func (c *Cargo) Connect() error {
  log.Println("Setup")
  s := New()
  log.Println("About to run")
  s.Run(8081)
  return nil
  // state, _ := dockercntrl.New()
  // err := state.VolumeCreateIdempotent("cargo")
  // return err
}
