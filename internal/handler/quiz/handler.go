package quiz

import (
	"diana-quiz/internal/service/auth"
	"diana-quiz/internal/service/quiz"
	"diana-quiz/internal/service/users"

	"github.com/danielblagy/go-utils/logger"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	logger   logger.Logger
	app      *fiber.App
	validate *validator.Validate

	quizService  quiz.Service
	usersService users.Service
	authService  auth.Service
}

func NewHandler(
	logger logger.Logger,
	app *fiber.App,
	validate *validator.Validate,
	quizService quiz.Service,
	usersService users.Service,
	authService auth.Service,
) *handler {
	h := &handler{
		logger:       logger,
		app:          app,
		validate:     validate,
		quizService:  quizService,
		usersService: usersService,
		authService:  authService,
	}
	h.setupRoutes()

	return h
}

func (h *handler) setupRoutes() {
	v1Group := h.app.Group("v1")

	testGroup := v1Group.Group("test")
	// GET test/ping - тестовый хендлер
	testGroup.Get("/ping", h.ping)

	usersGroup := v1Group.Group("users")
	// POST users/ - создать пользователя (регистрация)
	usersGroup.Post("/", h.createUser)
	// POST users/login - вход пользователя (генерация JWT токенов, токены записываются в http-only куки)
	usersGroup.Post("/login", h.logIn)
	// POST users/logout - выход пользователя (JWT токены в http-only куках удаляются) ТРЕБУЕТ АВТОРИЗАЦИИ
	usersGroup.Post("/logout", h.logOut)
	// GET users/quizes - получить список опросов пользователя ТРЕБУЕТ АВТОРИЗАЦИИ
	usersGroup.Get("/quizes", h.getUserQuizes)

	quizAdminGroup := v1Group.Group("quiz")
	// POST quiz/create - создать опрос ТРЕБУЕТ АВТОРИЗАЦИИ
	quizAdminGroup.Post("/create", h.createQuiz)
	// GET quiz/:id - получить опрос
	quizAdminGroup.Get("/:id", h.getQuiz)
	// POST quiz/:id/complete - пройти опрос ТРЕБУЕТ АВТОРИЗАЦИИ
	quizAdminGroup.Post("/:id/complete", h.completeQuiz)
	// GET quiz/:id/results - получить результаты опроса, доступно только владельцу опроса ТРЕБУЕТ АВТОРИЗАЦИИ
	quizAdminGroup.Get("/:id/results", h.getQuizResults)
}
