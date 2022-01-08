FROM golang:1.17 as builder

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /halo-discord-bot

FROM scratch

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

COPY --from=builder /halo-discord-bot /halo-discord-bot

CMD ["/halo-discord-bot"]