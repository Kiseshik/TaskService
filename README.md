# Task Service API (Go)

HTTP API для управления длительными I/O bound задачами на Go.

## Установка
```bash
git clone https://github.com/Kiseshik/TaskService.git
cd TaskService
go mod download
```
## Запуск сервера
```bash
go run cmd/main.go
```

## API Endpoints

| Метод  | Эндпоинт      | Описание               |
|--------|---------------|------------------------|
| `POST` | `/tasks`      | Создать новую задачу  |
| `GET`  | `/tasks/{id}` | Получить статус задачи | 
| `DELETE`| `/tasks/{id}`| Удалить задачу        |

## Архитектура
```bash
.
├── cmd/                 # Точка входа
├── pkg/
│   ├── handlers/        # HTTP обработчики
│   ├── models/          # Модели данных
│   ├── services/        # Бизнес-логика
│   └── storage/         # Хранение (in-memory)
├── go.mod               # Зависимости
├── go.sum               # Зависимости
```

## cURL-запросы для тестирования API 
```bash
# 1. Создать новую задачу
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json"

# 2. Получить статус задачи (замените {task_id} на реальный ID)
curl http://localhost:8080/tasks/{task_id} -H "Accept: application/json"

# 3. Удалить задачу (замените {task_id} на реальный ID)
curl -X DELETE http://localhost:8080/tasks/{task_id}
```

