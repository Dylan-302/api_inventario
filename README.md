# 🛒 API Inventario

API REST desarrollada en Go para la gestión de productos y categorías en un sistema de inventario.

---

## 🚀 Descripción

Esta API permite administrar un sistema de inventario mediante operaciones CRUD sobre productos y categorías. Está diseñada para ser rápida, escalable y fácil de integrar con aplicaciones frontend o móviles.

---

## 🧰 Tecnologías utilizadas

* Go (Golang)
* Gin (framework HTTP)
* GORM (ORM)
* PostgreSQL
* Variables de entorno (.env)

---

## 📁 Estructura del proyecto

```bash
.
├── db/
│   └── connection.go   # conexión a base de datos
├── models/
│   ├── product.go      # modelo de productos
│   └── category.go     # modelo de categorías
├── main.go             # punto de entrada
├── go.mod
├── go.sum
└── .env
```

---

## ⚙️ Configuración del entorno

1. Clonar el repositorio:

```bash
git clone https://github.com/Dylan-302/api_inventario.git
cd api_inventario
```

2. Crear archivo `.env`:

```env
DB_HOST=localhost
DB_USER=tu_usuario
DB_PASSWORD=tu_password
DB_NAME=tu_base_de_datos
DB_PORT=5432
```

3. Instalar dependencias:

```bash
go mod tidy
```

4. Ejecutar el proyecto:

```bash
go run main.go
```

---

## 📡 Endpoints

### 📦 Productos

* GET /products → Obtener todos los productos
* GET /products/:id → Obtener producto por ID
* POST /products → Crear producto
* PUT /products/:id → Actualizar producto
* DELETE /products/:id → Eliminar producto

---

### 🗂️ Categorías

* GET /categories → Obtener categorías
* POST /categories → Crear categoría

---

## 🧠 Modelo de datos

### Product

* ID
* Name
* Price
* CategoryID

### Category

* ID
* Name

---

## 🔐 Seguridad

Actualmente no implementa autenticación.
Se recomienda agregar JWT o middleware de seguridad para producción.

---

## 📌 Notas

* Proyecto en desarrollo
* Pensado como backend para e-commerce o sistema administrativo

---

## 👨‍💻 Autor

Dylan Barrios

---

## 📜 Licencia

Uso libre para fines educativos y proyectos personales.
