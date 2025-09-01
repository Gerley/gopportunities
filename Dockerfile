FROM golang:1.24-alpine

WORKDIR /app

# Instalar dependências
# RUN apk add --no-cache curl ca-certificates

# Instalar o migrate CLI
# ENV MIGRATE_VERSION=4.17.1

# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz \
#     | tar -xz && \
#     mv migrate /usr/local/bin/migrate && \
#     chmod +x /usr/local/bin/migrate


# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# RUN go build -o main ./cmd/server

EXPOSE 8080

# CMD ["./main"]
# CMD ['bash']
