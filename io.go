package cargo

import (
  "io/ioutil"
)

func get(filename string) ([]byte, error) {
  return ioutil.ReadFile(filename)
}


func put(filename string, data []byte) error {
  return ioutil.WriteFile(filename, data, 0644)
}
