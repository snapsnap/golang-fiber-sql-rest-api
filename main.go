package main

import (
	"fmt"
	"os"
	"rest-project/app/config"
	"rest-project/app/config/connections"
	"rest-project/app/routes"
	"rest-project/cmd"

	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

func main() {
	cnf := config.Get()
	db := connections.GetDatabase(cnf.Database)
	defer db.Close()

	// Menjalankan Migration atau Seeder berdasarkan argumen CLI
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			cmd.RunMigrations(db)
			fmt.Println("Migration completed!")
			return
		case "seed":
			cmd.RunSeeder(db)
			fmt.Println("Seeding completed!")
			return
		default:
			fmt.Println("Command not recognized")
			return
		}
	}

	app := fiber.New()

	file, loggerTest := config.RegisterLog(app)
	defer file.Close()

	routes.SetRouter(app, db)

	loggerTest.Println("Fiber server is running on : " + cnf.Server.Port)
	fiberlog.Info(app.Listen(cnf.Server.Host + ":" + cnf.Server.Port))

}
