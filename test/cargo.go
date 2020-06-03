package main

import (
  "github.com/armadanet/cargo"
  "log"
  "os"
)

func main() {
  // r1 := &cargo.Request{
  //   ReqType: 1,
  //   Name: "test.txt",
  // }
  r2 := &cargo.Request{
    ReqType: 2,
    Name: "test.txt",
    Data: []byte("This is a different string\n"),
  }
  r3 := &cargo.Request{
    ReqType: 1,
    Name: "test.txt",
  }
  requester := cargo.NewRequester()
  // resp1 := requester.SendRequest(r1)
  // log.Println(resp1.Status)
  // s1 := string(resp1.Data)
  // log.Println(s1)

  resp2 := requester.SendRequest(r2)
  log.Println(resp2.Status)
  s2 := string(resp2.Data)
  log.Println(s2)

  resp3 := requester.SendRequest(r3)
  log.Println(resp3.Status)
  s3 := string(resp3.Data)
  log.Println(s3)
  os.Exit(0)
}
