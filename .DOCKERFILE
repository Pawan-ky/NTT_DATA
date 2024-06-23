# Start with a base Go image
FROM golang:1.17-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build the Go app
RUN go build -o ntt_data

# Start a new stage from scratch
FROM alpine:latest  
WORKDIR /root/
COPY .env ./
COPY --from=builder /app/ntt_data .
# Command to run the executable
CMD ["./ntt_data"]
