FROM golang
WORKDIR /app
ADD go.mod /app/
RUN go mod download
ADD . /app/
RUN go build -o main cargo.go
ENTRYPOINT ["./main"]
