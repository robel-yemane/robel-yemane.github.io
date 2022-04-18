FROM golang:1.15 AS builder

WORKDIR /go/src/github.com/robel-yemane/automate

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd/ cmd
COPY articlewriter/ articlewriter
COPY articlereader/ articlereader

RUN go test ./... -cover

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o articlegen cmd/articlegen/articlegen.go

#FROM alpine:latest

#RUN apk --no-cache add ca-certificates
FROM busybox:glibc
#WORKDIR /root/
COPY --from=builder /go/src/github.com/robel-yemane/automate/articlegen /usr/bin/articlegen

ENTRYPOINT ["articlegen"]