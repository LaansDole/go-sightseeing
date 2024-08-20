# Temporal 101

This guide provides instructions for starting a Temporal server in development mode and initiating a simple workflow using the Temporal CLI.

## Prerequisites

Before running the commands in this guide, make sure you have the following installed:

- **Temporal CLI**: Command-line interface for interacting with Temporal.
- **Temporal Server**: The Temporal server should be set up locally.

## Starting the Temporal Server in Development Mode

To start the Temporal server in development mode, run the following command:

```bash
temporal server start-dev
```

### Description

- **Purpose**: This command starts a local instance of the Temporal server configured for development. It is designed for testing and development purposes and provides an easy way to run the Temporal server without needing a complex production setup.
- **Use Case**: Use this command when you want to develop and test Temporal workflows locally.

### Expected Output

Once started, the Temporal server should output logs indicating that it is running and listening for incoming workflow tasks. You can access the Temporal Web UI (usually at `http://localhost:8088`) to monitor workflows and view their status.

## Starting a Workflow

After the Temporal server is running, you can start a workflow using the following command:

```bash
temporal workflow start \
    --type GreetSomeone \
    --task-queue greeting-tasks \
    --workflow-id my-first-workflow \
    --input '"Donna"'
```

### Description

- **Purpose**: This command starts a new instance of the `GreetSomeone` workflow.
- **Options**:
    - `--type GreetSomeone`: Specifies the workflow type (name) to start. In this example, it's `GreetSomeone`.
    - `--task-queue greeting-tasks`: Specifies the task queue that the workflow worker listens to. The workflow tasks will be routed to this queue.
    - `--workflow-id my-first-workflow`: A unique identifier for this workflow instance. This ID is used to track and manage the workflow.
    - `--input '"Donna"'`: The input to the workflow. In this case, it passes the string `"Donna"` as an argument to the workflow.

### Example

This workflow could be designed to greet a person by name. The workflow logic might look something like:

```go
func GreetSomeone(ctx workflow.Context, name string) (string, error) {
    return fmt.Sprintf("Hello, %s!", name), nil
}
```
