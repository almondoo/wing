// Sample stdlogging writes log.Logger logs to the Cloud Logging.
package context

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/logging"
)

func Logging() {
	ctx := context.Background()

	// Creates a client.
	client, err := logging.NewClient(ctx, os.Getenv("PRODUCT_ID"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the log to write to.
	logName := "First-Art"

	logger := client.Logger(logName).StandardLogger(logging.Info)

	// Logs "hello world", log entry is visible at
	// Cloud Logs.
	logger.Println("Logging Start")
}
