package application

import (
	"context"
	"log"
	"os"
	"os/signal"
	"photo-upload-service/cmd/config"
	photoHandler "photo-upload-service/internal/handlers/photo"
	"photo-upload-service/internal/pkg/api/photo"
	"photo-upload-service/internal/rabbitmq"
	"photo-upload-service/internal/server"
	photoSrv "photo-upload-service/internal/service/photo"
	"syscall"
)

type App struct {
	cfg     *config.Cfg
	closers *closers
	srv     *server.Server
}

func New() *App {

	// init config
	cfg := &config.Cfg{}
	err := cfg.Init()
	mustInit("config", err)

	//append closer for log flush
	closers := &closers{}

	return &App{
		closers: closers,
		cfg:     cfg,
	}
}

func (a *App) Run(ctx context.Context) error {
	//graceful shutdown
	go a.signalHandler(ctx)

	//pgPool, err := db.InitPostgres(ctx, a.cfg.Database)
	//mustInit("postgres", err)
	//a.AddCloser(func() error { pgPool.Close(); return nil })

	//queue
	rabbit, err := rabbitmq.NewPublisher(a.cfg)
	if err != nil {
		return err
	}

	//services
	photoService := photoSrv.NewPhotoService(rabbit)

	//handlers
	photoConnector := photoHandler.NewPhotoHandler(photoService)

	//http server
	a.srv = server.New(
		a.cfg.Service.MainPort,
		a.cfg.Service.DebugPort,
	)

	photo.RegisterHandlers(a.srv.GetMainRouter(), photoConnector)

	return a.srv.Run(ctx)
}

func (a *App) AddCloser(c func() error) {
	a.closers.AddCloser(c)
}

func (a *App) Shutdown(_ context.Context) error {
	a.closers.Close()
	return nil
}

func (a *App) signalHandler(ctx context.Context) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	sig := <-ch
	log.Printf("os signal received: %s", sig)

	if err := a.Shutdown(ctx); err != nil {
		log.Printf("error app shutdown: %s", err)
	}
}

func mustInit(desc string, err error) {
	if err != nil {
		log.Fatalf("error to init %s: %s", desc, err)
	}
}
