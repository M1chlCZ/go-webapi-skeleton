package pkg

import (
	"context"
	"exbitron_info_app/pkg/database"
	"exbitron_info_app/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunAPI() {
	app := fiber.New(fiber.Config{
		AppName:       "Exbitron Info API",
		StrictRouting: false,
		WriteTimeout:  time.Second * 35,
		ReadTimeout:   time.Second * 35,
		IdleTimeout:   time.Second * 65,
	})
	app.Use(cors.New())
	app.Use(helmet.New(
		helmet.Config{
			ContentSecurityPolicy: "default-src 'self'",
		}))
	database.InitMySQL()
	utils.ReportMessage(fmt.Sprintf("EXBITRON API STARTED ON PORT 6900 | Version: %s", utils.VERSION))
	app.Get("/ping", ping)

	go func() {
		err := app.Listen(":6900")
		if err != nil {
			utils.WrapErrorLog(err.Error())
			panic(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	<-c
	_, cancel := context.WithTimeout(context.Background(), time.Second*15)
	utils.ReportMessage("/// = = Shutting down = = ///")
	defer cancel()
	_ = app.Shutdown()
	os.Exit(0)

}

func ping(c *fiber.Ctx) error {
	c.Set("Content-Security-Policy", "connect-src http://localhost:8080")
	return c.Status(200).SendString(fmt.Sprintf("Exbitron Info API | %s", utils.VERSION))
}
