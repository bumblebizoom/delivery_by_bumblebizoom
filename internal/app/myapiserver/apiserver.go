package myapiserver

import (
	logger "github.com/bumblebizoom/bumblogger"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type APIserver struct {
	config *Config
	router *mux.Router
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	logger.SetLogLevel(s.config.LogLevel)
	logger.DebugMsg("App logger was started in debug mod")
	return nil

}

func (s *APIserver) configureRouter() error {
	s.router.HandleFunc("/hallo", s.handleHallo())
	return nil
}

func (s *APIserver) handleHallo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hallo")
	}
}
