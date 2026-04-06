package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"photo-upload-service/internal/config"
	"photo-upload-service/internal/middleware"
	"time"
)

type Server struct {
	mainRouter  *gin.Engine
	mainHTTPSrv *http.Server
	mainPort    string
}

func New(cfg *config.Service) *Server {
	gin.SetMode(gin.DebugMode)

	mainRouter := gin.New()

	mainRouter.Use(middleware.RequestIDMiddleware())
	mainRouter.Use(gin.Recovery())

	mainRouter.GET("/health", func(c *gin.Context) {
		log.Println("Healthcheck request")
		c.Status(http.StatusOK)
	})

	return &Server{
		mainRouter:  mainRouter,
		mainHTTPSrv: newHTTPServer(cfg, mainRouter),
		mainPort:    cfg.MainPort,
	}
}

func newHTTPServer(cfg *config.Service, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.MainPort),
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}

func (s *Server) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		log.Printf("main server start and listen: %s", s.mainPort)
		if err := s.mainHTTPSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("error to start main server: %s", err)
		}

		return nil
	})

	return g.Wait()
}

func (s *Server) Shutdown(ctx context.Context, timeout time.Duration) error {
	//время на завершение запросов
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Printf("shutdown main server: %s", s.mainPort)
		if err := s.mainHTTPSrv.Shutdown(ctx); err != nil {
			return fmt.Errorf("error to shutdown main server: %s", err)
		}

		return nil
	})

	return g.Wait()
}

func (s *Server) GetMainRouter() *gin.Engine {
	return s.mainRouter
}
