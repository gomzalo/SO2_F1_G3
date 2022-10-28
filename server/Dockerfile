FROM golang:1.16.5 as developement

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/cespare/reflex@latest

ENV MONGO_USER Bases2
ENV MONGO_PASS SiSale

EXPOSE 8080

CMD reflex -g '*.go' go run main.go --start-service