# syntax=docker/dockerfile:1
FROM golang:1.22
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./client

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=0 /build/client ./client
CMD ["/app/client"]
