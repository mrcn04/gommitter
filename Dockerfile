# syntax=docker/dockerfile:1

FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

ARG GH_KEY
ARG GH_USERNAME
ARG REPO_NAME

ENV GH_KEY=${GH_KEY}
ENV GH_USERNAME=${GH_USERNAME}
ENV REPO_NAME=${REPO_NAME}

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /gommitter-app

EXPOSE 8080

# Run
CMD ["/gommitter-app"]