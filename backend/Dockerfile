FROM ghcr.io/golang/go:1.22

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]
