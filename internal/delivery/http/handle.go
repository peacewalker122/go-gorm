package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-gorm/domain/usecase"
	"go-gorm/internal/repository"
	uc "go-gorm/internal/usecase"
	"go-gorm/pkg/jwt"
	"go-gorm/pkg/middleware"
	"go-gorm/pkg/util"
	"net/http"
	"os"
)

var (
	ErrBadRequest = echo.NewHTTPError(http.StatusBadRequest, "bad request")
	ErrorNotFound = echo.NewHTTPError(http.StatusNotFound, "not found")
	ErrorInternal = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	ErrorConflict = echo.NewHTTPError(http.StatusConflict, "conflict")
	Unauthorized  = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")

	validate = validator.New()

	// db setup
	gormDriver   = util.InitializeGorm()
	userRepo     = repository.NewUserRepository(gormDriver)
	roomRepo     = repository.NewRoomRepository(gormDriver)
	ticketRepo   = repository.NewTicketRepository(gormDriver)
	userRoomRepo = repository.NewUserRoomRepository(gormDriver)
	// jwt setup
	key   = os.Getenv("SECRET_KEY")
	token = jwt.NewJwt(key)

	usecases = uc.NewUsecase(token, userRepo, roomRepo, ticketRepo, userRoomRepo, util.UUIDGenerator)
)

type Handler struct {
	usecase.UserUsecase
	usecase.TicketUsecase
	usecase.RoomUsecase
	echo  *echo.Echo
	Token jwt.Jwt
}

func InitHandler() *echo.Echo {
	var handler Handler

	handler.TicketUsecase = usecases
	handler.RoomUsecase = usecases
	handler.UserUsecase = usecases
	handler.Token = token

	echo := router(token, &handler)
	handler.echo = echo
	return echo
}

func router(jwtMethod jwt.Jwt, handle *Handler) *echo.Echo {
	e := echo.New()
	e.Validator = middleware.NewValidator(validate)
	e.HTTPErrorHandler = middleware.HTTPErrorHandler
	//e.Binder = &middleware.CustomBinder{}
	e.Use(
		middleware.MiddlewareLogging,
	)

	base := e.Group("/api/v1")

	userBase := base.Group("/user")
	userBase.POST("/register", handle.CreateUser)
	userBase.POST("/login", handle.Login)
	userBase.GET("/profile/:id", handle.GetUserProfile)

	authBase := base.Group("", middleware.AuthMiddleware(jwtMethod))

	roomBase := authBase.Group("/room")
	roomBase.POST("/create", handle.CreateRoom)
	roomBase.GET("/get/:id", handle.GetRoom)

	ticketBase := authBase.Group("/ticket")
	ticketBase.POST("/create", handle.CreateTicket)
	ticketBase.GET("/get/:id", handle.GetTicketByID)
	ticketBase.GET("/get-all", handle.GetAllTickets)

	return e
}
