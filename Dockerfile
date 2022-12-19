# First pull Golang image
FROM golang:1.19 AS builder
  
# Create diretory to application
WORKDIR /app

# Build application
COPY . .
 
# Run Stage
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server 

# multi-stage-build
FROM scratch
COPY --from=builder /app/server /server

ENTRYPOINT [ "/server" ]