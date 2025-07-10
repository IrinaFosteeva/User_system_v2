package main

import (
	"User_system_v2/internal/db"
	"User_system_v2/internal/handlers"
	"User_system_v2/internal/storage"
	"User_system_v2/internal/storage/memorystorage"
	"User_system_v2/internal/storage/postgresstorage"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env файл не найден, продолжаем с переменными окружения")
	}

	driver := os.Getenv("STORAGE_DRIVER")
	var store storage.Storage

	switch driver {
	case "memory":
		fmt.Println("🚀 Используем in-memory хранилище")
		store = memorystorage.NewMemoryStorage()

	case "postgres":
		fmt.Println("🐘 Используем PostgreSQL хранилище")
		db.Init()
		store = postgresstorage.NewPostgresStorage(db.DB)

	default:
		log.Fatalf("❌ Неизвестный STORAGE_DRIVER: %s", driver)
	}

	http.HandleFunc("/persons", handlers.MakePersonsHandler(store))
	http.HandleFunc("/persons/", handlers.MakePersonHandler(store))

	fmt.Println("✅ Сервер запущен на http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
