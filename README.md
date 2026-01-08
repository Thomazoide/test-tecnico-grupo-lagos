# Test técnico — Grupo Lagos

Proyecto dividido en backend (Go) y frontend (React + Vite + Tailwind).

## Backend (Go)
- Servidor HTTP en `:8080` con rutas:
	- `GET /products/page/{page}`: lista productos paginados desde OpenFoodFacts.
	- `GET /products/code/{code}`: obtiene un producto por código (OpenFoodFacts).
	- `POST /users`, `GET /users/{id}`, `PUT /cart/{id}`: endpoints de usuarios/carrito (requieren DB).
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
	 go build main.go
     ./main.exe
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

## Uso de IA
Para lo que mas la use fue para que me ayudara a crear este README


