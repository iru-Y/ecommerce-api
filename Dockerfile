FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 39998

CMD ["./main"]

