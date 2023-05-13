FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /goapp ./cmd/app

EXPOSE 8000

CMD ["/goapp"]