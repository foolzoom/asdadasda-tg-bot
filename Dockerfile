FROM golang:alpine
RUN go version

ADD . /go/src/app
WORKDIR /go/src/app

# Install git
RUN apk add --no-cache git
# Fetch deps
RUN go get
# Remove git
RUN apk del git

# Compile app
RUN go build -o main .
# Run app
CMD ["/go/src/app/main"]