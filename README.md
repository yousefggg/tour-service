# 🏔️ Tour Service (Mountain Tour Project)

**Tour Service** — это backend-сервис для системы бронирования горных туристических туров. Проект реализован на языке **Go** с соблюдением принципов **Clean Architecture** и микросервисного подхода.

## 📌 Описание

Сервис обеспечивает выполнение следующих задач:
*   **Управление турами**: полный цикл CRUD операций.
*   **Бронирование**: функционал заказа туров пользователями.
*   **Безопасность**: авторизация на базе JWT-токенов.
*   **Хранение данных**: использование PostgreSQL с поддержкой миграций.
*   **Интерфейс**: REST API для взаимодействия с фронтендом.

## ⚙️ Технологический стек

*   **Language:** Go (Golang)
*   **Database:** PostgreSQL (драйвер `pgx/v5`)
*   **Migrations:** `golang-migrate`
*   **Auth:** JWT (алгоритм HS256)
*   **Router:** Chi router
*   **Logging:** Zap logger
*   **Architecture:** Clean Architecture

---

## 🧱 Архитектура проекта

Проект разделен на логические блоки для упрощения тестирования и поддержки:

*   `internal/`
    *   `domain`: Бизнес-модели и сущности.
    *   `usecase`: Слой бизнес-логики.
    *   `repository`: Слой работы с базой данных (Persistence).
    *   `delivery`: Слой доставки (HTTP-хендлеры, роутинг, middleware).
*   `common-lib/`: Общие модули (логгер, JWT-утилиты, конфиги, обработка ошибок).

---

## 🚀 Запуск проекта

### 1. Клонирование репозитория
```bash
git clone https://github.com/yousefggg/tour-service.git
cd tour-service
```

### 2. Настройка окружения
Создайте файл `.env` в корне проекта и заполните его:
```env
APP_PORT=8081
APP_ENVIRONMENT=local
APP_LOG_LEVEL=debug

DATABASE_URL=postgres://user:password@localhost:5432/tour_db?sslmode=disable
DATABASE_MAX_OPEN_CONNS=10
DATABASE_MAX_IDLE_CONNS=5
DATABASE_CONN_TIMEOUT=5s

AUTH_JWT_SECRET=your-secret-key
AUTH_TOKEN_TIME=24h
```

### 3. Запуск PostgreSQL через Docker
```bash
docker run --name tour-postgres \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=tour_db \
  -p 5432:5432 \
  -d postgres
```

### 4. Старт сервиса
Миграции применяются автоматически при запуске приложения:
```bash
go run main.go
```

---

## 📡 API Endpoints

### 🔓 Публичные методы (Public)


| Метод | Эндпоинт | Описание |
| :--- | :--- | :--- |
| `GET` | `/api/v1/tours` | Получить список всех туров |
| `GET` | `/api/v1/tours/{id}` | Получить детальную информацию о туре |

### 🔐 Защищенные методы (JWT required)
**Бронирования (Bookings):**


| Метод | Эндпоинт | Описание |
| :--- | :--- | :--- |
| `POST` | `/api/v1/bookings` | Создать новое бронирование |
| `GET` | `/api/v1/bookings` | Список бронирований текущего пользователя |
| `GET` | `/api/v1/bookings/{id}` | Детали конкретного бронирования |

**Администрирование туров (Admin):**


| Метод | Эндпоинт | Описание |
| :--- | :--- | :--- |
| `POST` | `/api/v1/admin/tours` | Создать новый тур |
| `PUT` | `/api/v1/admin/tours/{id}` | Обновить данные тура |
| `DELETE` | `/api/v1/admin/tours/{id}` | Удалить тур из системы |

---

## 📦 Модель бронирования (JSON)

```json
{
  "tour_id": "uuid",
  "phone_number": "+123456789",
  "people_count": 2,
  "notes": "хочу номер с видом",
  "medical_info": "нет противопоказаний",
  "payment_method": "card"
}
```

---

## 🧪 Тестирование и качество

*   **Запуск тестов:** `go test ./... -v -cover`
*   **Покрытие (Coverage):** ~60% (основной упор на слой `usecase`).

### Особенности реализации:
*   Полное разделение бизнес-логики и транспортного слоя.
*   Внедрение зависимостей (**Dependency Injection**) через конструкторы.
*   Использование **Mocks** для Unit-тестирования.

---

## 🏗️ Roadmap (Планы по развитию)

- [ ] Интеграция Swagger/OpenAPI документации.
- [ ] Интеграционные тесты с использованием `testcontainers`.
- [ ] Механизм Refresh Tokens.
- [ ] Ролевая модель доступа (RBAC: Admin/User).
- [ ] Наблюдаемость: метрики (Prometheus) и трассировка.