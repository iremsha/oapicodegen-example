package app

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/iremsha/oapicodegen-example/internal/config"
	"github.com/iremsha/oapicodegen-example/internal/database"
	"github.com/iremsha/oapicodegen-example/internal/entity"
	bankgen "github.com/iremsha/oapicodegen-example/internal/gen/bank"
	cardgen "github.com/iremsha/oapicodegen-example/internal/gen/card"
	"github.com/iremsha/oapicodegen-example/internal/handler"
	logger "github.com/iremsha/oapicodegen-example/internal/log"
	"github.com/iremsha/oapicodegen-example/internal/repository"
	"github.com/iremsha/oapicodegen-example/internal/service"
	"go.elastic.co/apm/module/apmfiber/v2"
)

const migrationsPath = "migrations"

type App struct {
	cfg config.Config
	log *logger.Logger
}

func New(cfg config.Config, log *logger.Logger) *App {
	return &App{cfg, log}
}

func (app *App) Run() error {
	fiberApp := fiber.New(fiber.Config{
		ReadTimeout:           time.Second * time.Duration(app.cfg.App.ReadTimeout),
		ReadBufferSize:        app.cfg.App.ReadBufferSize,
		UnescapePath:          true,
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ErrorHandler:          app.errorHandler,
	})

	app.setupMiddlewares(fiberApp)

	if err := database.MigrateUp(app.cfg.Database, app.log, migrationsPath); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	database, err := database.New(app.cfg, app.log)
	if err != nil {
		return fmt.Errorf("database initialization failed: %w", err)
	}

	app.registerHandlers(fiberApp, database)

	go app.gracefulShutdown(fiberApp)

	return fiberApp.Listen(fmt.Sprintf(":%s", app.cfg.App.Port))
}

func (app *App) setupMiddlewares(fiberApp *fiber.App) {
	prometheus := fiberprometheus.New("oapicodegen-example")
	prometheus.RegisterAt(fiberApp, "/metrics")
	fiberApp.Use(prometheus.Middleware)

	grp := fiberApp.Group("/api")
	grp.Use(app.authMiddleware())
	grp.Use(apmfiber.Middleware())
}

func (app *App) errorHandler(ctx *fiber.Ctx, err error) error {
	errMsg := err.Error()
	app.log.Error(ctx.Context(), errMsg)
	return ctx.Status(fiber.StatusInternalServerError).JSON(entity.ResponseJSON{
		Success: false,
		Message: errMsg,
	})
}

func (app *App) authMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if string(c.Request().Header.Peek(app.cfg.App.Credentials.Header)) != app.cfg.App.Credentials.APIKey {
			app.log.Warn(c.Context(), "Unauthorized access attempt")
			return c.Status(fiber.StatusUnauthorized).JSON(entity.ResponseJSON{
				Success: false,
				Message: "Unauthorized",
			})
		}
		return c.Next()
	}
}

func (app *App) gracefulShutdown(fiberApp *fiber.App) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-exit
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := fiberApp.ShutdownWithContext(timeoutCtx); err != nil {
		app.log.Error(context.Background(), "Graceful shutdown failed: "+err.Error())
	}
}

func (app *App) registerHandlers(fiberApp *fiber.App, database *database.Database) {
	cardRepository := repository.NewCardRepository(database.DB)
	cardService := service.NewCardService(cardRepository)
	cardHandler := handler.NewCardHandler(cardService)

	cardgen.RegisterHandlers(fiberApp, cardHandler)

	bankRepository := repository.NewBankRepository(database.DB)
	bankService := service.NewBankService(bankRepository)
	bankHandler := handler.NewBankHandler(bankService)

	bankgen.RegisterHandlers(fiberApp, bankHandler)
}
