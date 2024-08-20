package api

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
)

func main() {}

func GreetSomeone(ctx workflow.Context, name string) (string, error) {
	return fmt.Sprintf("Hello " + name + "!"), nil
}
