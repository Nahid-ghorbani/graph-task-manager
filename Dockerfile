FROM golang:1.25.1

WORKDIR /graph-task-manager

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

# Install curl to check Postgres readiness
RUN apt-get update && apt-get install -y curl netcat-openbsd && rm -rf /var/lib/apt/lists/*
