package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

const (
	readTimeout  = 10 * time.Second
	writeTimeout = 10 * time.Second
	idleTimeout  = 15 * time.Second

	shutdownTimeout = 3 * time.Second
)

type Server struct {
	mainRouter   *gin.Engine
	debugRouter  *gin.Engine
	mainHTTPSrv  *http.Server
	debugHTTPSrv *http.Server
	mainPort     string
	debugPort    string
}

func New(mainPort, debugPort string) *Server {
	gin.SetMode(gin.DebugMode)

	mainRouter := gin.New()

	mainRouter.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	debugRouter := gin.New()

	return &Server{
		mainRouter:   mainRouter,
		mainHTTPSrv:  newHTTPServer(buildLocalAddr(mainPort), mainRouter),
		mainPort:     mainPort,
		debugRouter:  debugRouter,
		debugHTTPSrv: newHTTPServer(buildLocalAddr(debugPort), debugRouter),
		debugPort:    debugPort,
	}
}

func newHTTPServer(addr string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
}

func (s *Server) Run(ctx context.Context) error {

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		log.Printf("debug server start and listen: %s", s.debugPort)
		if err := s.debugHTTPSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("error to start debug server: %s", err)
		}

		return nil
	})

	g.Go(func() error {
		log.Printf("main server start and listen: %s", s.mainPort)
		if err := s.mainHTTPSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("error to start main server: %s", err)
		}

		return nil
	})

	return g.Wait()
}

func (s *Server) Shutdown(ctx context.Context) error {
	//время на завершение запросов
	ctx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		log.Printf("shutdown debug server: %s", s.debugPort)
		if err := s.debugHTTPSrv.Shutdown(ctx); err != nil {
			return fmt.Errorf("error to shutdown debug server: %s", err)
		}

		return nil
	})

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

func buildLocalAddr(port string) string {
	return fmt.Sprintf(":%s", port)
}
