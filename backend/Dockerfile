# Etapa de compilación / Desarrollo
FROM golang:1.24-alpine AS dev

# Instalamos git por si alguna dependencia lo necesita
RUN apk add --no-cache git

WORKDIR /app

# Copiar dependencias primero
COPY go.mod ./
# COPY go.sum ./
RUN go mod download

# Copiamos el resto del código
COPY . .

RUN rm -rf vendor

# Instalamos 'air' para tener Hot Reload (recarga automática al guardar)
# Esto es opcional pero recomendado para proyectos personales
RUN go install github.com/cosmtrek/air@v1.51.0

# Exponemos el puerto
EXPOSE 4000

# Por defecto usamos air para desarrollo, o go run si prefieres lo simple
CMD ["air"]