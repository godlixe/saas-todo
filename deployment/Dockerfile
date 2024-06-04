FROM golang:1.21.1-alpine

WORKDIR /usr/src/app

# RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o /build/app" -command="/build/app"

# RUN go mod download

# CMD ["air", "-c", ".air.toml"]
