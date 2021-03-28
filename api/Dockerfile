# build stage
FROM golang:1.15-buster as builder

WORKDIR /app

COPY . .

RUN go mod download
# RUN install -v -mod=vendor ./cmd/...

RUN go build -o main ./

# Run stage
FROM debian:buster

RUN apt-get update --fix-missing && \
apt-get install -yqq ca-certificates

COPY --from=builder /app .


# RUN chmod +x ./main

# EXPOSE 8080
CMD ["./main"]