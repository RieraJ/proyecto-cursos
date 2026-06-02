# --- Etapa de Desarrollo (hot reload con CRA dev server) ---
FROM node:22-alpine3.21 AS dev

WORKDIR /app

COPY package.json pnpm-lock.yaml* ./
RUN corepack enable pnpm && pnpm i --ignore-scripts

# Copia el código para que la imagen funcione sin volumen también
COPY . .

EXPOSE 3000

# Necesario para que el file watcher de webpack detecte cambios dentro de Docker
ENV CHOKIDAR_USEPOLLING=true
ENV WATCHPACK_POLLING=true
# Hace que el WebSocket de HMR use el mismo puerto que la página (3000)
ENV WDS_SOCKET_PORT=0

CMD ["pnpm", "start"]

# --- Etapa de Construcción (Build) ---
FROM node:22-alpine3.21 AS build

WORKDIR /app

COPY package.json pnpm-lock.yaml* ./
RUN corepack enable pnpm && pnpm i --ignore-scripts

COPY . .
RUN pnpm run build

# --- Etapa de Producción (Servidor Web) ---
FROM nginx:1.27-alpine AS production
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
