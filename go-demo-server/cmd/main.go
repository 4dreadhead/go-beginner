package main

import (
	"fmt"
	"goadvancedserver/configs"
	"goadvancedserver/internal/auth"
	"goadvancedserver/internal/link"
	"goadvancedserver/internal/user"
	"goadvancedserver/pkg/db"
	"goadvancedserver/pkg/middleware"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()
	database := db.NewDB(config)

	linkRepo := link.NewRepository(database)
	userRepo := user.NewRepository(database)

	authService := auth.Service{UserRepository: userRepo}

	host := ":8081"
	auth.NewAuthHandler(router, auth.HandlerDependencies{Config: config, Service: &authService})
	link.NewLinkHandler(router, link.HandlerDeps{Config: config, Repository: linkRepo})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	server := http.Server{
		Addr:    host,
		Handler: stack(router),
	}

	fmt.Printf("Listening on %v\n", host)
	server.ListenAndServe()
}
