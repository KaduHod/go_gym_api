FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go_gym_api

# Fase 2: Node modules + assets
FROM node:20-slim AS node_modules_builder
WORKDIR /app

# Copia apenas o que é necessário para instalar dependências
COPY package.json package-lock.json ./
RUN npm install --omit=dev

# Fase 3: Imagem final
FROM debian:bookworm-slim
WORKDIR /app

# Copia o binário Go
COPY --from=builder /app/go_gym_api /app/go_gym_api
# Após FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
# Instala certificados CA na imagem final

# Copia variáveis de ambiente
COPY .env /app/.env
COPY .env.prod /app/.env.prod
COPY views /app/views
COPY public /app/public
RUN mkdir -p /app/logs && \
    touch /app/logs/server.log && \
    chmod 666 /app/logs/server.log

# Copia os assets públicos
COPY public /app/public

# Copia os node_modules instalados
COPY --from=node_modules_builder /app/node_modules /app/node_modules
COPY package.json /app/package.json

# Expõe a porta da aplicação
EXPOSE 3005

# Executa o binário Go
CMD ["/app/go_gym_api", "prod"]

