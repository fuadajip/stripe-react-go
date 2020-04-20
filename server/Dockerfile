## Builder
FROM golang:1.11-alpine as build_base
RUN apk update && apk upgrade && \
    apk --no-cache --update add bash git make gcc g++ libc-dev

WORKDIR /go/src/github.com/fuadajip/stripe-react-go/server

ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

#This is the ‘magic’ step that will download all the dependencies that are specified in 
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download 
# command will _ only_ be re-run when the go.mod or go.sum file change 
# (or when we add another docker instruction this line)
RUN go mod download

# This image builds the weavaite server
FROM build_base AS server_builder
# Here we copy the rest of the source code
WORKDIR /go/src/github.com/fuadajip/stripe-react-go/server
COPY . .
RUN ls -lh
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN make engine

## #In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --no-cache --update add ca-certificates tzdata && \
    mkdir /stripe-react-go && mkdir /server && mkdir /app
WORKDIR /stripe-react-go/server/app

EXPOSE 12001

COPY --from=server_builder /go/src/github.com/fuadajip/stripe-react-go/server/engine .
COPY --from=server_builder /go/src/github.com/fuadajip/stripe-react-go/server/app.config* ./
RUN ls -lh
CMD /stripe-react-go/server/app/engine