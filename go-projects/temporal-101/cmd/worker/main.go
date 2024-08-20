package main

import (
	"github.com/laansdole/go-sightseeing/go-projects/temporal-101/cmd/api"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "greeting-tasks", worker.Options{})

	w.RegisterWorkflow(api.GreetSomeone)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
