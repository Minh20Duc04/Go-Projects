package main

import (
	"log"
	"net/http"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Minh20Duc04/Go-Projects/internal/products"
)

type application struct {
	config config
	//logger
}

type config struct {
	addr string //8080
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middleware stack chuẩn
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Health check endpoints (rất quan trọng cho Kubernetes, Docker, monitoring)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"OK","timestamp":"` + time.Now().Format(time.RFC3339) + `"}`))
	})

	// Nhiều team còn tách health vs ready
	r.Get("/ready", func(w http.ResponseWriter, r *http.Request) {
		// Sau này bạn sẽ check DB connection ở đây
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("READY"))
	})

	// Route gốc để test nhanh
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to E-commerce API - Chi version"))
	})

	// Các route API thật của bạn sẽ nằm đây sau
	r.Route("/api/v1", func(r chi.Router) {
		h := products.NewHandler(nil)
		r.Get("/products", h.ListProducts)
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at addr : %s", srv.Addr)

	return srv.ListenAndServe()
}
