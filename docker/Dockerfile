# Etapa de build
FROM golang:1.24 AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos do projeto para o container
COPY . .

# Baixa as dependências
RUN go mod tidy

# Compila o binário
RUN go build -o server .

# Etapa final
FROM debian:bookworm-slim

# Instala dependências mínimas (certificados TLS)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Diretório de trabalho
WORKDIR /app

# Copia o binário da etapa anterior
COPY --from=builder /app/server .

# Copia imagens (ajuste se for necessário)
COPY --from=builder /app/images ./images

# Define variável de ambiente padrão (pode ser sobrescrita com -e)
ENV LISTEN_ADDRESS=8080

# Expõe a porta definida
EXPOSE 8080

# Comando de entrada
CMD ["./server"]