package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// RegisterLog mengatur logging aplikasi
func RegisterLog(app *fiber.App) (*os.File, *log.Logger) {
	// Buat folder logs jika belum ada
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			log.Fatalf("Error creating logs folder: %v", err)
		}
	}

	// Logging ke file per hari
	logFileName := fmt.Sprintf("logs/log_%s.log", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	// Middleware logger untuk Fiber
	app.Use(logger.New(logger.Config{
		Output: file,
		Format: "[${time}] ${status} - ${method} ${path} | ${latency}\n",
	}))

	// Logger tambahan untuk mencetak log manual
	loggerTest := log.New(file, "", log.LstdFlags)
	loggerTest.Println("Logger initialized...")

	return file, loggerTest
}
