package main

import (
	"context"
	"fmt"
	"os"

	"diana-quiz/internal/config"
	"diana-quiz/internal/db"
	"diana-quiz/internal/handler/quiz"
	"diana-quiz/internal/service/auth"
	quiz_service "diana-quiz/internal/service/quiz"
	"diana-quiz/internal/service/users"

	config_wrapper "github.com/danielblagy/go-utils/config-wrapper"
	"github.com/danielblagy/go-utils/logger"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	logger := logger.NewLogger()
	logger.InfoKV("Hello from diana quiz!")

	// load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		logger.FatalKV("can't load .env file", "err", err.Error())
		os.Exit(1)
	}

	ctx := context.Background()

	// connect to postgres database
	pgConnectionPool, err := pgxpool.New(ctx, config_wrapper.GetEnvValue(config.DatabaseUrl).String())
	if err != nil {
		logger.FatalKV("can't connect to database", "err", err.Error())
		os.Exit(1)
	}
	defer pgConnectionPool.Close()

	// fiber app

	app := fiber.New()
	app.Use(fiberLogger.New())

	// validator
	validate := validator.New()

	// db
	queryFactory := db.NewQueryFactory(pgConnectionPool)

	// services

	quizService := quiz_service.NewService(logger, queryFactory)
	usersService := users.NewService(pgConnectionPool)
	authService := auth.NewService(usersService)

	// handlers
	_ = quiz.NewHandler(
		logger.AddContext("handler", "diana-quiz"),
		app,
		validate,
		quizService,
		usersService,
		authService,
	)

	if startAppErr := app.Listen(fmt.Sprintf(":%s", config_wrapper.GetEnvValue(config.ServerPort).String())); startAppErr != nil {
		logger.FatalKV("can't start fiber app", "err", startAppErr.Error())
	}
}
