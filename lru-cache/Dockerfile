# Copyright: 2024 Arm Ltd. All Rights Reserved.

# Build the app in the large Golang docker image which has all the
# compilers and tools needed to build the code.
FROM golang:alpine AS builder
ADD . /lru-cache
WORKDIR /lru-cache
RUN mkdir -p /build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/lru-cache main.go

# Copy the built binary to a much more lightweight image to reduce
# the size of the finished image.
FROM alpine
COPY --from=builder /build/* /app/
WORKDIR /app
# Create service user
RUN addgroup -S fib && adduser -S fib -G fib -u 1001
RUN chown 1001 /app
USER 1001
EXPOSE 8081
ENTRYPOINT [ "./lru-cache" ]
