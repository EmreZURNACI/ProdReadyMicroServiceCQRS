package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	cmd_c "github.com/EmreZURNACI/ProdFullReadyApp_User/internal/controller/user/command"
	query_c "github.com/EmreZURNACI/ProdFullReadyApp_User/internal/controller/user/query"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/infra/mongo"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/infra/postgres"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/middleware"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/service/user/command"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/service/user/query"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func Start(pgHandler *postgres.Repository, mongoHandler *mongo.Repository, appconfig *config.Appconfig) {

	server := fiber.New(
		fiber.Config{
			Prefork:       true,
			ServerHeader:  "User Service",
			StrictRouting: false,
			CaseSensitive: false,
			BodyLimit:     1024 * 1024 * 5,
			Concurrency:   256 * 1024,
			AppName:       "ProdFullReadyApp",
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.JSON(fiber.Map{"error": err.Error()})
			},
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  10 * time.Second,
		},
	)

	server.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	v1 := server.Group("/api/v1")
	users := v1.Group("/users")
	users.Use(middleware.IsAuthorized)

	cmd_service := command.NewUserService(pgHandler)
	cmd_controller := cmd_c.NewUserHandler(cmd_service)

	query_service := query.NewUserService(mongoHandler)
	query_controller := query_c.NewUserHandler(query_service)

	users.Post("/", cmd_controller.Create)
	users.Put("/:id", cmd_controller.Update)
	users.Delete("/:id", cmd_controller.Delete)
	users.Get("/", query_controller.Get)
	users.Get("/:id", query_controller.GetID)

	zap.L().Info("app starting...", zap.Int("port", appconfig.SrvConfig.Port))
	if err := server.Listen(fmt.Sprintf(":%d", appconfig.SrvConfig.Port)); err != nil {
		zap.L().Error(err.Error())
		return
	}

	gracefulShutdown(server)
}
func gracefulShutdown(app *fiber.App) {
	//Create channel for shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	//Wait for shutdown signal
	<-sigChan
	zap.L().Info("Shutting down server...")

	//Shutdown with 5 second timeout
	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		zap.L().Error("Error during server shutdown", zap.Error(err))
	}

	zap.L().Info("Server gracefully stopped")
}
