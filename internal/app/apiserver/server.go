package apiserver

import (
	"net/http"

	"github.com/Dukastlik/avitomx-api.git/internal/app/products"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer Structure
type APIServer struct {
	config   *Config
	logger   *logrus.Logger
	router   *mux.Router
	products *products.Products
}

// New APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start APIserver
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	if err := s.configureProducts(); err != nil {
		return err
	}

	s.logger.Info("Starting server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// configureLogger sets logging level
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

// configureProducts
func (s *APIServer) configureProducts() error {
	pr := products.New(s.config.Products)
	if err := pr.Open(); err != nil {
		return err
	}
	s.products = pr
	return nil
}
