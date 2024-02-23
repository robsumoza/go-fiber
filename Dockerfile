FROM golang:1.22-alpine AS build

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod go.sum ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY *.go ./

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build ldflags="-s -w" -o /app .

FROM scratch

WORKDIR /

COPY --from=build /app /app

COPY --from=build /etc/passwd /etc/passwd

USER 1001

EXPOSE 8080

ENTRYPOINT ["/app"]
