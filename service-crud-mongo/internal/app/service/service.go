package service

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/xamust/quizRBI/service-crud-mongo/internal/app/cash"
	"github.com/xamust/quizRBI/service-crud-mongo/internal/app/store"
	"net/http"
	"sync"
)

type Service struct {
	config     *Config
	logger     *logrus.Logger
	store      *store.Store
	repository store.Repository
	cash       *cash.Cash
	ctx        context.Context
	mu         *sync.Mutex
	mux        *mux.Router
}

// init new server
func New(config *Config) *Service {
	return &Service{
		config: config,
		logger: logrus.New(),
		ctx:    context.Background(),
		mux:    mux.NewRouter(),
		mu:     &sync.Mutex{},
	}
}

// configure logrus...
func (s *Service) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

// config store MongoDB....
func (s *Service) configureStore() error {
	s.store = store.New(s.config.Store)
	if err := s.store.Open(); err != nil {
		return err
	}
	s.repository = s.store
	return nil
}

// config cash...
func (s *Service) configureCash() error {
	s.cash = cash.New(s.config.Cash, s.ctx, s.store, s.mu, s.logger)
	if err := s.cash.UpdateCash(); err != nil {
		return err
	}
	return nil
}

// config route...
func (s *Service) configureRouter() {
	s.mux.HandleFunc("/api/v1/data", s.requestHandler)
}

func (s *Service) StartService() error {

	if err := s.configureLogger(); err != nil {
		return fmt.Errorf("can't configure logger: %v", err)
	}

	if err := s.configureStore(); err != nil {
		s.logger.Errorf("can't configure store: %v", err)
		return fmt.Errorf("can't configure store: %v", err)
	}
	go func(service *Service) error {
		err := service.configureCash()
		if err != nil {
			service.logger.Errorf("can't configure cash: %v", err)
			return err
		}
		return nil
	}(s)

	s.configureRouter()
	s.logger.Info(fmt.Sprintf("Starting server (bind on %v)...", s.config.BindAddr))
	return http.ListenAndServe(s.config.BindAddr, s.mux)
}
