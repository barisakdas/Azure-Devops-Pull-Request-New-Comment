FROM golang:alpine
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .
Run cp /build/Config/config.json .
CMD ["/dist/main"]