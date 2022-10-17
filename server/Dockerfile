FROM golang:bullseye

WORKDIR /go/src/app
COPY . .

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . . 

RUN go build -o /go/bin/app cmd/main.go

EXPOSE 8080

CMD ["app"]