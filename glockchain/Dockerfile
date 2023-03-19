ARG GO_VERSION=1.20.2

FROM golang:${GO_VERSION}-alpine as builder
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app ./cmd/app/

FROM scratch
COPY --from=builder /etc/group /etc/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/app /app

USER appuser

ENTRYPOINT ["/app"]