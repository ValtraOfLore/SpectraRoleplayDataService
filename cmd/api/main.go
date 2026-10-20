package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ValtraOfLore/SpectraRoleplayDataService/internal/config"
	guilddata "github.com/ValtraOfLore/SpectraRoleplayDataService/internal/guild-data"
	"github.com/ValtraOfLore/SpectraRoleplayDataService/internal/health"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// Configuration Instantiation
	envPtr := flag.String("env", "local", "The environment this project is run on. Valid values: local, production")
	flag.Parse()

	fmt.Println("Environment set to: " + *envPtr)

	config, configErr := config.Load(*envPtr)
	if configErr != nil {
		fmt.Println("Error loading configs")
		return
	}
	fmt.Println("Configurations loaded: ")
	fmt.Printf("%+v\n", *config)

	// Database Connection
	dbClient, dbConnErr := mongo.Connect(options.Client().ApplyURI(config.DatabaseConnectionString))
	if dbConnErr != nil {
		fmt.Println("Unable to connect to database")
		return
	}
	guildColl := dbClient.Database("bot_data").Collection("guild")

	mux := http.NewServeMux()

	// Services
	healthService := health.NewService()
	guildService := guilddata.NewService(guildColl)

	// Handlers
	healthHandler := health.NewHandler(healthService)
	guildHandler := guilddata.NewHandler(guildService)

	// Routes: health
	mux.HandleFunc("GET /health", healthHandler.Get)

	// Routes: guild-info
	mux.HandleFunc("GET /guild-data/{guild_id}", guildHandler.Get)

	http.ListenAndServe(":"+config.Port, mux)
}
