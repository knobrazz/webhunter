FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o web-hunter .

FROM alpine:latest

# Install required tools
RUN apk add --no-cache \
    nmap \
    masscan \
    git \
    python3 \
    py3-pip \
    && pip3 install nuclei

WORKDIR /app
COPY --from=builder /app/web-hunter .

ENTRYPOINT ["./web-hunter"]

