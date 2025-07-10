# Используем официальный образ Golang
FROM golang:1.24

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN go build -o main .

# Указываем порт
EXPOSE 8080

# Команда запуска
CMD ["./main"]
