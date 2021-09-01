FROM golang:1.17-alpine AS build
#Install git
RUN apk add --no-cache git
#Get the hello world package from a GitHub repository
RUN go env -w GOPROXY=direct
# Clear GOPATH for go.mod
ENV GOPATH=
# cache dependencies
ADD go.mod go.sum ./
RUN go mod download
# Build the project and send the output to /bin/HelloWorld
ADD . .
RUN go build -o /bin/sandbox-go-http

FROM golang:1.17-alpine
#Copy the build's output binary from the previous build container
COPY --from=build /bin/sandbox-go-http /bin/sandbox-go-http
ENTRYPOINT ["/bin/sandbox-go-http"]
