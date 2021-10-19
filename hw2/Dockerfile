FROM golang:1.16-alpine AS build

COPY src /go/src/project/

WORKDIR /go/src/project/

RUN go mod init && go build -o /bin/httpserver_example
ENTRYPOINT ["/bin/httpserver_example"]

# FROM scratch

# COPY --from=build /bin/httpserver_example /bin/httpserver_example

# ENTRYPOINT ["/bin/httpserver_example"]