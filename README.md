# 👤 User System (v2)

Простое REST API на Go для управления пользователями (persons), с использованием PostgreSQL, Docker и миграций.

---

## 🚀 Функциональность

- ✅ Создание пользователя
- 📋 Получение всех пользователей
- 🔍 Получение пользователя по ID
- ✏️ Обновление данных пользователя
- 🗑️ Удаление пользователя

---

## 🛠️ Стек технологий

- Go
- PostgreSQL
- Docker & Docker Compose
- [migrate](https://github.com/golang-migrate/migrate) — для миграций

---

## ⚙️ Установка и запуск

1. **Клонируйте репозиторий:**

```bash
git clone https://github.com/IrinaFosteeva/User_system_v2.git
cd User_system_v2 
```

2. **Создайте .env файл:**

```bash
POSTGRES_USER=postgres
POSTGRES_PASSWORD=123postgres
POSTGRES_DB=myappdb
POSTGRES_PORT=5432
```

3. **Запустите проект через Docker:**

```bash
docker-compose up --build -d
```

# Примеры API-запросов:

Получить всех пользователей:

GET http://localhost:8080/persons

Создать пользователя:

POST http://localhost:8080/persons
Content-Type: application/json

{
"name": "Alice",
"age": 30
}

Получить по ID:

GET http://localhost:8080/persons/1


Обновить пользователя:

PATCH http://localhost:8080/persons/1
Content-Type: application/json

{
"name": "Alice Smith"
}

Удалить пользователя:

DELETE http://localhost:8080/persons/1

