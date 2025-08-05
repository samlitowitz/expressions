############################
# -- go-test
############################
FROM golang:1.24-alpine AS go-test

WORKDIR $GOPATH/src/github.com/samlitowitz/expressions
COPY . .

ENV GO111MODULE=on
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go mod download
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go mod verify

CMD [ "go", "test", "-v", "./test-cases/..." ]
