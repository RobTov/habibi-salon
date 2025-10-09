package main

import (
	"log"
	"net/http"
	"time"

	"github.com/RobTov/habibi-salon/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	address string
	db      dbConfig
	env     string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))
	// CORS
	r.Use(cors.Handler(cors.Options{
		// TODO: change this for env variables
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowedOriginFunc: func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           30,
	}))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		r.Route("/appointments", func(r chi.Router) {
			r.Route("/", func(r chi.Router) {
				r.Get("/", app.getAppointmentHandler)
				r.Post("/", app.createAppointmentHandler)
			})
		})

		r.Route("/clients", func(r chi.Router) {
			r.Route("/", func(r chi.Router) {
				r.Get("/", app.getClientHandler)
				r.Post("/", app.createClientHandler)
			})

			r.Route("/{clientID}", func(r chi.Router) {
				r.Use(app.clientsContextMiddleware)
				r.Get("/", app.getClientByIDHandler)
				r.Patch("/", app.updateClientHandler)
				r.Delete("/", app.deleteClientHandler)
			})
		})

		r.Route("/services", func(r chi.Router) {
			r.Route("/", func(r chi.Router) {
				r.Get("/", app.getServiceHandler)
				r.Post("/", app.createServiceHandler)
			})

			r.Route("/{serviceID}", func(r chi.Router) {
				r.Use(app.servicesContextMiddleware)
				r.Get("/", app.getServiceByIDHandler)
				r.Patch("/", app.updateServiceHandler)
				r.Delete("/", app.deleteServiceHandler)
			})
		})

		r.Route("/stock", func(r chi.Router) {
			r.Route("/", func(r chi.Router) {
				r.Get("/", app.getStockHandler)
				r.Post("/", app.createStockHandler)
			})

			r.Route("/{stockID}", func(r chi.Router) {
				r.Use(app.stockContextMiddleware)
				r.Get("/", app.getStockByIDHandler)
				r.Patch("/", app.updateStockHandler)
				r.Delete("/", app.deleteStockHandler)
			})
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server running at: http://localhost%s\n", app.config.address)
	return srv.ListenAndServe()
}
