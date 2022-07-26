FROM golang:alpine

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o server ./server.go

EXPOSE 8080

CMD [ "./server" ]
