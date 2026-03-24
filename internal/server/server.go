package server

import (
	"database/sql"
	"ims/internal/handler"
	"ims/internal/repository"
	"ims/internal/route"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultServerPort = ":8080"

func InitHTTPServer(db *sql.DB) *http.Server {

	productRepo := repository.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepo)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	route.SetupRoutes(router, productHandler)

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultServerPort
	}

	server := &http.Server{
		Addr:    serverPort,
		Handler: router,
	}

	return server
}
