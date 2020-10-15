FROM golang:1.14

RUN mkdir -p /app/go-app-sample
WORKDIR /app/go-app-sample
COPY go.mod .
COPY . .
EXPOSE 8080

RUN go build -o app
CMD ["./app"]