
# Fase 1: Builder
FROM golang:1.23 AS builder
WORKDIR /app

# Copia os arquivos do projeto e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte e compila o binário estático
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go_gym_api

# Fase 2: Imagem final
FROM debian:bookworm-slim
WORKDIR /app

# Copia apenas o binário compilado e os arquivos necessários
COPY --from=builder /app/go_gym_api /app/go_gym_api
COPY .env /app/.env

# Expõe a porta da aplicação
EXPOSE 3004

# Define o comando de execução
CMD ["/app/go_gym_api"]

