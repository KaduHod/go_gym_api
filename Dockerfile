
FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go_gym_api
# Fase 1: imagem base para desenvolvimento com Air
FROM golang:1.23

# Instala o Air e outras dependências
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copia os arquivos de dependência primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Expõe a porta
EXPOSE 3005

# Comando para rodar com Air (live reload)
CMD ["air"]

