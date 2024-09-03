package main

import (
	"context"
	"fmt"
)

func main() {}

// runTemporalWorker init and runs the Temporal Worker in a blocking fashion
// Worker will be shutdown gracefully when ctx.Done() is closed

func runTemporalWorker(ctx context.Context, cfg Configuration) error {
	// No need to change

	auth, err := jwtauth.NewAuthenticator(ctx, &cfg.Auth)
	if err != nil {
	}

	temporalHostPort := fmt.Sprintf("%s:%d", cfg.Temporal.Host, cfg.Temporal.Port)

	temporalClient, err := temporalconn.WorkerDialCluster(ctx, temporalHostPort)
	if err != nil {
	}

	log.Info()
	defer temporalClient.Close()

	melbourne, err := autime.LoadLocationMelbourne
	// May need to refactor

}
