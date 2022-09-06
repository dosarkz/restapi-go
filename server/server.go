package server

import (
	"context"
	"github.com/rs/cors"
	"gorm.io/gorm"
	"rest/app/helpers/env"
	"rest/app/validators"
	"rest/config"
	"rest/config/initial"
	"rest/database/connections"
	"rest/database/migrations"
	"rest/database/redis"
	"rest/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "rest/router/lang"
)

func Run() {
	l := log.New(os.Stdout, "rest-api\n", log.LstdFlags)
	handler := router.GetRouter()

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ORIGIN_ALLOWED")},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	s := &http.Server{
		Addr:         config.GetApp().GetUrlAddr(), // configure the bind address
		Handler:      c.Handler(handler),           // the default handlers
		IdleTimeout:  120 * time.Second,            // max time for connections using TCP keep-alive
		ReadTimeout:  20 * time.Second,             // max time to read request from client
		WriteTimeout: 30 * time.Second,             // max time to write response to the client
	}
	log.Printf("Server connected to http://localhost:%s/ for mux", config.GetApp().Port)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt) //nolint:govet
	signal.Notify(sigChan, os.Kill)      //nolint:govet

	sig := <-sigChan
	l.Println("Terminate received, gracefully shuttingdown...", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer func(database *gorm.DB) {
		db, err := database.DB()
		if err != nil {
			log.Fatal(err)
		}
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}(connections.GetDB())
	err := s.Shutdown(tc)
	if err != nil {
		return
	}
}

func LoadDependencies() {
	dbConf := new(initial.Database)
	env.HasEnvironment()
	config.Init()
	connections.LoadDB(dbConf)
	redis.ConnectToRedis(new(redis.Config))
	migrations.Migrate(dbConf)
	validators.Init(connections.GetDB())
	initial.LoadRepos(connections.GetDB())
	initial.LoadServices(initial.GetRepositories())
	initial.LoadControllers(initial.GetServices(), connections.GetDB())
}
