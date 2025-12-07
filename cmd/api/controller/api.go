package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/ariefzainuri96/go-api-ecommerce-api-gateway/cmd/api/docs"
	"github.com/ariefzainuri96/go-api-ecommerce-api-gateway/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-ecommerce-api-gateway/grpc"
	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Application struct {
	Config    Config
	Service   *grpc.ServerService
	Validator *validator.Validate
}

type Config struct {
	Addr string
}

func (app *Application) Mount() *http.ServeMux {
	mux := http.NewServeMux()

	// mux.Handle("/v1/product/", middleware.Authentication(http.StripPrefix("/v1/product", app.ProductRouter())))

	// mux.Handle("/v1/cart/", middleware.Authentication(http.StripPrefix("/v1/cart", app.CartRouter())))

	// mux.Handle("/v1/order/", middleware.Authentication(http.StripPrefix("/v1/order", app.OrderRouter())))

	mux.Handle("/v1/auth/", http.StripPrefix("/v1/auth", app.AuthRouter()))

	// register health check
	mux.Handle("GET /health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		jsonResponse, _ := json.Marshal(map[string]string{
			"status": "OK",
		})
		w.Write(jsonResponse)
	}))

	// mux.Handle("/v1/xendit-callback/", http.StripPrefix("/v1/xendit-callback", app.XenditCallbackRouter()))

	mux.Handle("/v1/swagger/", httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	))

	return mux
}

func (app *Application) Run(mux *http.ServeMux) error {

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Recoverer,
	)

	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      stack(mux),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	log.Printf("Server has started on %s", app.Config.Addr)

	return srv.ListenAndServe()
}
