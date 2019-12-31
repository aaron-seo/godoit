# build stage
FROM golang:latest AS builder

WORKDIR /app
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o GoDoIt 

# final stage
FROM alpine:latest AS production

WORKDIR /root/
COPY --from=builder /app .

EXPOSE 8080

CMD ["./GoDoIt"]
