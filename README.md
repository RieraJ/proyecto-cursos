# Proyecto Cursos

Aplicación full-stack para la gestión de cursos.

## Estructura del repositorio

- `backend/`: API REST escrita en Go. Gestiona usuarios, cursos, inscripciones y comentarios utilizando el framework Gin y la base de datos MySQL.
  - `app/`: configuración del router y mapeo de rutas.
  - `clients/`: conexión a la base de datos y operaciones CRUD.
  - `controllers/`: lógica de manejo de peticiones HTTP.
  - `dao/` y `dto/`: modelos de datos y objetos de transferencia.
  - `middleware/`: autenticación y autorización.
  - `services/`: capa de servicios para la lógica de negocio.
  - `main.go`: punto de entrada que carga variables de entorno, conecta con la base de datos y arranca el servidor.
- `frontend/`: aplicación React creada con Create React App.
- `docker-compose.yml`: orquestación de MySQL, la API y el frontend.
- `backend/init.sql`: script para crear la base de datos inicial.

## Ejecución con Docker

1. Tener instalado Docker y Docker Compose.
2. Ejecutar `docker-compose up --build` en la raíz del proyecto.
3. La API quedará disponible en `http://localhost:4000` y el frontend en `http://localhost:3000`.

## Pruebas

- Backend: `go test ./...`
- Frontend: `npm test` dentro de `frontend/`.

