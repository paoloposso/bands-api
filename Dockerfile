FROM golang:1.15-alpine

WORKDIR $GOPATH/src/github.com/github.com/paoloposso/bands-auth-api

ENV MONGO_TIMEOUT=120
ENV PORT=80

COPY go.mod .
COPY go.sum .

COPY . .

RUN go mod tidy

RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 80

# Run the executable
CMD ["github.com/paoloposso/bands-auth-api"]