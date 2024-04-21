package main

import (
	"context"
	"database/sql"
	"github.com/SawitProRecruitment/UserService/middlewares"
	"github.com/SawitProRecruitment/UserService/modules"
	"github.com/SawitProRecruitment/UserService/services"
	"os"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dbDsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbDsn)

	if err != nil {
		panic(err)
	}

	dbConn, err := db.Conn(context.Background())

	if err != nil {
		panic(err)
	}

	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(dbConn)

	var repo repository.UserRepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Conn: dbConn,
	})

	svc := initServices(repo)
	initMiddlewares(e, svc)

	var server generated.ServerInterface = newServer(svc)

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func initServices(repo repository.UserRepositoryInterface) services.Services {

	passwordAuth := modules.BcryptPasswordAuth{}

	userService := services.NewUserService(repo, passwordAuth)

	privateKey, err := os.ReadFile("cert/id_rsa")

	if err != nil {
		panic("cannot read private key")
	}

	publicKey, err := os.ReadFile("cert/id_rsa.pub")

	if err != nil {
		panic("cannot read public key")
	}

	jwtAuth := modules.NewRS256Jwt(privateKey, publicKey)

	authenticationService := services.NewAuthenticationService(repo, passwordAuth, jwtAuth)

	return services.Services{
		Authentication: authenticationService,
		User:           userService,
	}
}

func initMiddlewares(e *echo.Echo, svc services.Services) {

	verifyJwtMiddleware := middlewares.NewVerifyJwtMiddleware(svc)

	e.Use(verifyJwtMiddleware.Process)
}

func newServer(svc services.Services) *handler.Server {

	opts := handler.NewServerOptions{
		UserService:           svc.User,
		AuthenticationService: svc.Authentication,
	}
	return handler.NewServer(opts)
}
