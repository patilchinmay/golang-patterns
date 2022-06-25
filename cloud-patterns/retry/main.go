// https://github.com/avast/retry-go
// https://stackoverflow.com/questions/62900451/how-to-customize-http-client-or-http-transport-in-go-to-retry-after-timeout

package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/avast/retry-go"
)

// Effector is an type. Any function that needs to be retried should conform to the effector type. e.g. simulateTransientError conforms to the Effector type.
type Effector func() (string, error)

var count int

func simulateTransientError() (string, error) {
	count++
	fmt.Printf("Received request: %d\n", count)

	if count <= 3 {
		return "intentional fail", errors.New("intentional fail error")
	} else {
		return "success", nil
	}
}

func fetchWithRetries(effector Effector) (message string, err error) {
	retry.Do(
		// The actual function that does "stuff"
		func() error {
			message, err = effector()
			return err
		},
		retry.OnRetry(func(try uint, orig error) {
			fmt.Printf("Request %d failed. Retrying #: %d\n", try+1, try+2)
		}),
		retry.Attempts(5),
		// Basically, we are setting up a delay
		// which randoms between 0 and 2 seconds.
		retry.Delay(1*time.Second),
		retry.MaxJitter(1*time.Second),
	)

	return
}

func main() {
	message, err := fetchWithRetries(simulateTransientError)

	if err != nil {
		fmt.Printf("Error fetching data: %s", err)
		os.Exit(1)
	}

	fmt.Println("message : ", message)
}
