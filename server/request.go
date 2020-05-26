package server

type ReqType int

const (
  ReqType Read = 1
  ReqType Write = 2
)

type Request struct {
  ReqType
}
