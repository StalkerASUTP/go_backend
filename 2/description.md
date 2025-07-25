# Программа обработки данных - Задание 2

## Описание
Консольная программа на Go, которая выполняет обработку числовых данных из JSON файла и осуществляет HTTP запросы с логированием всех операций.

## Функциональность
Программа выполняет следующие операции:

1. **Чтение данных из JSON файла** - считывает массив целых чисел из файла `json/data.json`
2. **Вычисление суммы** - подсчитывает общую сумму всех чисел в массиве
3. **HTTP запрос** - выполняет GET запрос на указанный URL и проверяет код ответа
4. **Логирование** - записывает результаты каждого шага в файл `logs/data.log`

## Архитектура
Программа построена с использованием принципов чистой архитектуры:

- **Factory Pattern** - для создания компонентов системы
- **Interface Segregation** - четкое разделение интерфейсов для различных компонентов
- **Dependency Injection** - внедрение зависимостей через конструкторы
- **Structured Logging** - использование slog для структурированного логирования

## Структура проекта
```
2/
├── cmd/main.go                    # Точка входа в приложение
├── internal/
│   ├── calculator/calculator.go   # Калькулятор для подсчета суммы
│   ├── config/config.go          # Загрузка конфигурации
│   ├── factory/factory.go        # Фабрика компонентов
│   ├── http/client.go            # HTTP клиент
│   ├── interfaces/interfaces.go  # Интерфейсы компонентов
│   ├── logger/struct_logger.go   # Структурированный логгер
│   ├── reader/reader.go          # Чтение данных из файла
│   └── service/data_processor.go # Основная бизнес-логика
├── json/data.json                # Файл с исходными данными
├── logs/data.log                 # Файл логов
└── go.mod                        # Модуль Go
```

## Конфигурация
Конфигурация загружается из переменных окружения (файл `.env`):

- `URL` - URL для HTTP запроса (по умолчанию: https://example.com)
- `JSON` - путь к JSON файлу с данными (по умолчанию: json/data.json)
- `LOGS` - путь к файлу логов (по умолчанию: logs/data.log)

## Формат JSON файла
```json
[
    0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711
]
```

## Пример вывода
```
Total sum of numbers is: 46367
✅ SUCCESS: Host: example.com, Status: 200 OK
```

## Обработка ошибок
Программа корректно обрабатывает все возможные ошибки:
- Ошибки чтения файла
- Ошибки парсинга JSON
- Ошибки HTTP запросов
- Ошибки логирования

Все ошибки логируются с детальной информацией и программа завершается с соответствующим кодом выхода.