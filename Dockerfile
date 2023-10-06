FROM golang:1.21.2-alpine3.18 AS builder

LABEL maintainer="Renzo Beltran <rnzbeltran+dev@gmail.com>"

# Move to working directory (/build).
WORKDIR /go/src/api-resume-service

# Copy files for caching.
COPY go.mod go.sum /go/src/api-resume-service/

# Download dependecies for caching.
RUN go mod download

# Copy the code into the container.
COPY cmd /go/src/api-resume-service/cmd

COPY config /go/src/api-resume-service/config
COPY internal /go/src/api-resume-service/internal

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o /app ./cmd/api

FROM scratch

WORKDIR /go/src/api-resume-service
# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder /go/src/api-resume-service/config/ /go/src/api-resume-service/config
COPY --from=builder "/app" "/app"

EXPOSE 3000
# Command to run when starting the container.
ENTRYPOINT ["/app"]