package main

import (
	"context"
	"github.com/laansdole/go-sightseeing/go-projects/temporal-101/cmd/api"
	"github.com/sirupsen/logrus"
	"go.temporal.io/sdk/client"
	"log"
	"os"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "example-workflow",
		TaskQueue: "greetings-tasks",
	}

	w, err := c.ExecuteWorkflow(context.Background(), options, api.GreetSomeone, os.Args[1])
	if err != nil {
		logrus.Fatalln("unable to execute workflow", err)
	}
	logrus.Println("Started workflow", " WorkflowID ", w.GetID(), " RunID ", w.GetRunID())

	var result string
	err = w.Get(context.Background(), &result)
	if err != nil {
		logrus.Fatalln("unable to get result", err)
	}
	logrus.Println(result)
}
