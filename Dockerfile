FROM golang:1.24.0-alpine3.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o api main.go

# Stage 2: Run
FROM alpine:3.21
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /app/api .
EXPOSE 80
CMD [ "./api" ]
