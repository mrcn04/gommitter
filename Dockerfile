FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /gommitter-app

ENV PORT=8081
EXPOSE 8081

# Run
CMD ["/gommitter-app"]