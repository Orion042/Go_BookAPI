FROM golang:1.21.5-alpine3.19

WORKDIR /api

COPY ./go.sum .

COPY ./go.mod .

RUN go mod download

COPY ./main.go .

CMD ["go", "run", "main.go"]