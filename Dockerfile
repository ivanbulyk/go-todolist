FROM golang:1.15-alpine3.12 as builder
WORKDIR $GOPATH/src/github.com/ivanbulyk/go-todolist
COPY ./ .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v
RUN cp go-todolist /

FROM alpine:latest AS production
COPY --from=builder / /
CMD ["/go-todolist"]
