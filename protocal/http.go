package protocal

import (
	"flag"
	"hugeman/configs"
	"hugeman/internal/handlers"
	"hugeman/internal/repositories"
	"hugeman/internal/services"
	"hugeman/pkg/database_driver/gorm"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type config struct {
	ENV string `mapstructure:"env"`
}

// ServeHTTP func
func ServeHTTP() error {
	app := fiber.New()
	var cfg config
	flag.StringVar(&cfg.ENV, "env", "", "the environment to use")
	flag.Parse()
	configs.InitViper("./configs", cfg.ENV)
	logrus.Info(configs.GetViper().Env)

	dbConGorm, err := gorm.ConnectToPostgreSQL(
		configs.GetViper().Postgres.Host,
		configs.GetViper().Postgres.Port,
		configs.GetViper().Postgres.Username,
		configs.GetViper().Postgres.Password,
		configs.GetViper().Postgres.DbName,
		configs.GetViper().Postgres.SSLMode,
	)
	if err != nil {
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("Gracefull shut down ...")
			gorm.DisconnectPostgres(dbConGorm.Postgres)
			err := app.Shutdown()
			if err != nil {
				log.Println("Error when shutdown server: ", err)
			}
		}
	}()

	postgresRepo := repositories.NewPostgres(dbConGorm.Postgres)
	srv := services.New(postgresRepo)
	hdl := handlers.New(srv, dbConGorm.Postgres)
	app.Get("/health", hdl.HealthCheck)

	hugeman := app.Group("/v1/api")
	{
		hugeman.Post("/todo", hdl.CreateTodo)
		hugeman.Put("/todo", hdl.UpdateTodo)
		hugeman.Delete("/todo/:id", hdl.DeleteTodo)
		hugeman.Get("/todo/:id", hdl.GetTodo)
		hugeman.Get("/todo", hdl.GetTodo)
	}

	err = app.Listen(":" + configs.GetViper().App.Port)
	if err != nil {
		return err
	}

	logrus.Println("Listerning on port: ", configs.GetViper().App.Port)
	return nil
}
