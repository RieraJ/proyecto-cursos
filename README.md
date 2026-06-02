# Proyecto Cursos

Plataforma full-stack para gestión de cursos online. Los usuarios pueden explorar el catálogo, inscribirse en cursos y dejar comentarios. Los administradores pueden crear, editar y eliminar cursos, y gestionar los roles de los usuarios.

## Stack

| Capa | Tecnologías |
|---|---|
| **Backend** | Go 1.24 · Gin · GORM · MySQL 8 |
| **Frontend** | React 18 · React Router v6 · pnpm |
| **Infra** | Docker · Docker Compose · nginx |

**Dependencias frontend destacadas:** `js-cookie` · `sweetalert2` · `lucide-react` · `react-icons` · `react-responsive`

---

## Quickstart

```bash
docker-compose up --build
```

| Servicio | URL |
|---|---|
| Frontend | http://localhost:3000 |
| API | http://localhost:4000 |
| MySQL | localhost:3307 |

---

## Variables de entorno

### Backend — `backend/.env`

```env
PORT=4000
SECRET=tu_secreto_jwt
DB=user:password@tcp(db:3306)/cursos
```

### Frontend

```env
REACT_APP_API_URL=http://localhost:4000
```

Si no se define `REACT_APP_API_URL`, el frontend apunta a `http://localhost:4000` por defecto.

---

## Desarrollo local (sin Docker)

### Backend

```bash
cd backend
go run main.go
# o con hot reload:
air
```

### Frontend

```bash
cd frontend
pnpm install
pnpm start
```

---

## Estructura del proyecto

```
.
├── backend/
│   ├── app/            # Router Gin y mapeo de rutas
│   ├── clients/        # Conexión a MySQL (GORM) y auto-migrate
│   ├── controllers/    # Handlers HTTP: parsean request y llaman al servicio
│   ├── dao/            # Structs GORM + queries a la BD
│   ├── dto/            # Shapes de request/response (separados de dao/)
│   ├── middleware/     # requireAuth (JWT) y requireAdmin
│   ├── services/       # Lógica de negocio
│   ├── init.sql        # Seed de datos iniciales
│   └── main.go         # Punto de entrada
└── frontend/
    └── src/
        ├── components/ # Componentes React, cada uno con su .css
        ├── App.jsx     # Definición de rutas (React Router)
        ├── ThemeContext.js  # Contexto para tema oscuro/claro
        ├── config.js   # API_URL desde variable de entorno
        └── utils.js    # Helpers: formatLength, validateImageFile
```

> El esquema de la base de datos es gestionado por GORM (`AutoMigrate`). `init.sql` solo inserta datos iniciales.

---

## Features

### Usuarios
- Registro e inicio de sesión con JWT almacenado en cookie
- Dos roles: **student** (por defecto) y **admin**
- Edición de perfil (nombre, contraseña)
- Carga y cambio de foto de perfil (base64)

### Cursos
- Catálogo de cursos con búsqueda por nombre
- Detalle de curso con categorías y duración formateada
- Inscripción a cursos
- Vista de cursos inscritos por usuario

### Comentarios
- Publicar y eliminar comentarios en cualquier curso
- Historial de comentarios por usuario

### Administración (solo admins)
- Crear, editar y eliminar cursos (con imagen en base64)
- Panel administrativo: listado de todos los usuarios y cursos
- Cambio de rol de usuario (student ↔ admin)

### UI
- Tema oscuro / claro con persistencia en `localStorage`
- Diseño responsive (`react-responsive`)

---

## API — Endpoints

### Autenticación y usuarios

| Método | Ruta | Auth | Descripción |
|---|---|---|---|
| `POST` | `/signup` | — | Registro |
| `POST` | `/login` | — | Login, devuelve cookie JWT |
| `POST` | `/logout` | — | Invalida la cookie |
| `GET` | `/user-info` | JWT | Info del usuario autenticado |
| `GET` | `/users` | JWT + Admin | Lista todos los usuarios |
| `PUT` | `/users/me` | JWT | Editar perfil propio |
| `PUT` | `/users/me/photo` | JWT | Cambiar foto de perfil |
| `PUT` | `/update-user-type` | JWT + Admin | Cambiar rol de un usuario |

### Cursos

| Método | Ruta | Auth | Descripción |
|---|---|---|---|
| `GET` | `/courses` | — | Listar todos los cursos |
| `GET` | `/search-courses?q=...` | — | Buscar cursos por nombre |
| `GET` | `/users/:id/courses` | JWT | Cursos inscritos de un usuario |
| `POST` | `/courses` | JWT + Admin | Crear curso |
| `PUT` | `/courses/:id` | JWT + Admin | Editar curso |
| `DELETE` | `/courses/:id` | JWT + Admin | Eliminar curso |

### Inscripciones y comentarios

| Método | Ruta | Auth | Descripción |
|---|---|---|---|
| `POST` | `/enroll` | JWT | Inscribirse en un curso |
| `GET` | `/courses/:id/comments` | JWT | Comentarios de un curso |
| `GET` | `/users/:id/comments` | JWT | Comentarios de un usuario |
| `POST` | `/comments` | JWT | Publicar comentario |
| `DELETE` | `/comments/:id` | JWT | Eliminar comentario |

---

## Tests

```bash
# Backend
cd backend && go test ./...

# Frontend
cd frontend && pnpm test
```
