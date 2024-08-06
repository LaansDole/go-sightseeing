package main

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	_defaultPort = "8080"
)

func main() {
	/* Listen for termination signals:
	* SIGINT handles Ctrl+C locally
	 */
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	startServerErrCh := make(chan error)
	mainCtx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context, errCh chan error) {
		errCh <- runAPIServer(ctx)
	}(mainCtx, startServerErrCh)

	select {
	case err := <-startServerErrCh:
		logrus.Error(mainCtx, err, "\nFailed to start server")
	case sig := <-signalCh:
		logrus.Info(
			mainCtx,
			"received signal: %v",
			sig.String(),
		)
		cancel()
	}
}

func runAPIServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	muxServer := &http.Server{
		ReadTimeout: 5 * time.Second,
		Addr:        ":" + _defaultPort,
		Handler:     mux,
	}

	errChan := make(chan error)
	go func(errCh chan error) {
		logrus.Info(ctx, "\nlistening on port: "+_defaultPort)
		if err := muxServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}(errChan)

	select {
	case err := <-errChan:
		logrus.Error(ctx, err, "\nFailed to ListenAndServe()")
		return err
	case <-ctx.Done():
		logrus.Warn(ctx, "\nShutting down server as ctx.Done() is closed")
		shutdownAPIServer(ctx, muxServer)
		return nil
	}
}

func shutdownAPIServer(ctx context.Context, server *http.Server) {
	shutdownTimeout := 10 * time.Second
	shutdownCtx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.Error(ctx, err, "\nfailed to shutdown API server: %v")
	}
	logrus.Warn(ctx, "\nShutting down server completed")
}
