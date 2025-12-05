// @title           Your Ecommerce API
// @version         1.0
// @description     This is the documentation for the main e-commerce service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /v1

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"fmt"
	"log"
	"os"

	ctrl "github.com/ariefzainuri96/go-api-ecommerce-api-gateway/cmd/api/controller"
	"github.com/ariefzainuri96/go-api-ecommerce-api-gateway/cmd/api/docs"
	"github.com/ariefzainuri96/go-api-ecommerce-api-gateway/grpc"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func buildServerService() (*grpc.ServerService, error) {
	authClient, err := grpc.NewAuthGRPCClient(os.Getenv("AUTH_SERVICE_ADDR"))
	if err != nil {
		return nil, fmt.Errorf("auth service: %w", err)
	}

	return grpc.NewServerService(authClient), nil
}

func main() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
			return
		}
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("%v:%v", os.Getenv("ADDR"), os.Getenv("PORT"))

	cfg := ctrl.Config{
		Addr: fmt.Sprintf(":%v", os.Getenv("PORT")),
	}

	validate := validator.New()

	serverService, err := buildServerService()

	if err != nil {
		log.Fatalf("failed to start gateway: %v", err)
	}

	app := &ctrl.Application{
		Config:    cfg,
		Service:   serverService,
		Validator: validate,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
