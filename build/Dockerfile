FROM golang
WORKDIR /app
ADD go.mod go.sum /app/
RUN go mod download
ADD . /app/
ARG location=cmd
RUN go build -o main ${location}/cargo.go
ENTRYPOINT ["./main"]
