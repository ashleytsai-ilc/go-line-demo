FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download && \
    go build -o gotoline

ENTRYPOINT ["./gotoline"]

EXPOSE 8080