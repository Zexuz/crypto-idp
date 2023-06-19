FROM golang:1.20.5-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o idp cmd/main.go

FROM alpine:latest

# Copy the binary from the first stage
COPY --from=0 /app/idp /idp
COPY ./assets app/assets

# Set the binary as the entrypoint of the container
ENTRYPOINT ["/idp"]