FROM golang:1.15-alpine

WORKDIR $GOPATH/src/github.com/paoloposso/bands-api

ENV MONGO_TIMEOUT=15
ENV PORT=80

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 80

# Run the executable
CMD ["bands-api"]