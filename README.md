# go-parcel-tracker 📦

Сервис для отслеживания посылок с использованием базы данных SQLite.

---

## Описание / Description

**RU**: REST-подобный сервис на Go для управления посылками: создание, изменение адреса доставки, удаление и получение информации о статусе. Данные хранятся в SQLite.

**EN**: A Go-based service for parcel tracking: create, update delivery address, delete, and retrieve parcel status. Data is stored in SQLite.

---

## Основные возможности / Features

- ➕ Создание новой посылки с указанием адреса
- ✏️ Изменение адреса доставки
- ❌ Удаление посылки
- 🔍 Получение информации о посылке по ID
- 📋 Список всех посылок

---

## Технологии / Tech Stack

- **Go** 1.21+
- **SQLite** — лёгкая встроенная БД
- **database/sql** — стандартный пакет для работы с БД
- **testing** + **assert** — юнит-тесты

---

## Установка и запуск / Installation

```bash
# Клонирование
git clone https://github.com/V-SemYo/go-parcel-tracker.git
cd go-parcel-tracker

# Установка зависимостей
go mod tidy

# Запуск
go run main.go parcel.go
Примечание: база данных tracker.db создаётся автоматически при первом запуске.

Структура проекта / Project Structure
text
go-parcel-tracker/
├── main.go           # Инициализация БД и основной цикл
├── parcel.go         # Логика работы с посылками (CRUD)
├── parcel_test.go    # Юнит-тесты с использованием assert
├── tracker.db        # Файл базы данных SQLite
├── go.mod
└── go.sum
API / Основные функции
go
// Создание посылки
func CreateParcel(address string) (int64, error)

// Получение посылки по ID
func GetParcel(id int64) (*Parcel, error)

// Обновление адреса
func UpdateAddress(id int64, newAddress string) error

// Удаление посылки
func DeleteParcel(id int64) error

// Получение всех посылок
func GetAllParcels() ([]Parcel, error)
Тестирование / Testing
Тесты покрывают основные операции и используют библиотеку assert для проверок.

Запуск тестов:

bash
go test -v ./...
Пример теста:

go
func TestCreateAndGetParcel(t *testing.T) {
    id, err := CreateParcel("Москва, ул. Тверская, 1")
    assert.NoError(t, err)
    
    parcel, err := GetParcel(id)
    assert.NoError(t, err)
    assert.Equal(t, "Москва, ул. Тверская, 1", parcel.Address)
}
Особенности реализации / Implementation Details
Оптимизированные SQL-запросы (использование одного запроса для SetAddress и Delete)

Обработка ошибок на всех уровнях

Автоматическое создание таблицы при первом запуске

Закрытие соединения с БД при завершении работы

