# Test técnico — Grupo Lagos

Proyecto dividido en backend (Go) y frontend (React + Vite + Tailwind).

## Backend (Go)
- Servidor HTTP en `:8080` con rutas:
	- `GET /products/page/{page}`: lista productos paginados desde OpenFoodFacts.
	- `GET /products/code/{code}`: obtiene un producto por código (OpenFoodFacts).
	- `GET /users`, `GET /users/{id}`, `PUT /cart/{id}`: endpoints de usuarios/carrito (requieren DB).
- Conexión a MySQL vía `.env` usando GORM y migraciones para `User` y `Cart`.

### Correr backend
1. Crear `.env`:
	 ```env
	 DBUSER=usuario
	 DBPASS=pass
	 DBHOST=127.0.0.1
	 DBPORT=3306
	 DBNAME=mi_base
	 ```
2. Ejecutar:
	 ```powershell
	 cd backend
	 go run .
	 ```
	 El servidor queda en `http://localhost:8080`.

## Frontend (React)
- Interfaz con búsqueda arriba:
	- Escribe para filtrar en vivo por nombre, marca o código.
	- Pulsa “Buscar” para consultar por código exacto al backend.
- Lista de productos en grilla responsiva y paginación al final (Primero/Anterior/Siguiente/Último).

### Correr frontend
```powershell
cd frontend
npm install
npm run dev
```
frontend: `http://localhost:5173`; backend: `http://localhost:8080`.

## Nota
No tuve tiempo de integrar todos los algoritmos solicitados debido a mi trabajo actual, pero fue entretenido avanzar y dejar una base funcional con paginación, filtrado en vivo y búsqueda por código.

## Ejemplos Postman
- Productos por página:
	1) Nuevo request GET
	2) URL: `http://localhost:8080/products/page/1`
	3) Headers: `Accept: application/json`
	4) Send. Respuesta envuelta en `ResponsePayload` con `data.products`.

- Producto por código:
	1) Nuevo request GET
	2) URL: `http://localhost:8080/products/code/{code}` (ej: `737628064502`)
	3) Headers: `Accept: application/json`
	4) Send. Respuesta envuelta en `ResponsePayload` con `data.product`.


