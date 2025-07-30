FROM golang:1.24

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o /app/app .

WORKDIR /app

EXPOSE 8080

CMD ["./app"]
