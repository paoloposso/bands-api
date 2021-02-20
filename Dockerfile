FROM golang:1.15-alpine

WORKDIR $GOPATH/src/github.com/paoloposso/bands-api

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Run the executable
CMD ["bands-api"]