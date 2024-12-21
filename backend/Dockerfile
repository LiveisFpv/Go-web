# Используем базовый образ с Go
FROM golang:1.23

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы зависимостей
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Копируем весь код backend
COPY backend ./

WORKDIR /app/main
RUN go build -o main .

# Указываем порт, который использует приложение
EXPOSE 15432
WORKDIR /app/main
# Команда для запуска
CMD ["./main"]
