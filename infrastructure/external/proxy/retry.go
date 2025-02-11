package proxy

import (
	"context"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v5"
)

// RetryService is a generic function signature for any service that may need retries
type RetryService func() (interface{}, error)

// ExecuteWithRetry applies retry logic with exp. backoff on a given service.
func ExecuteWithRetry(service RetryService) (interface{}, error) {

	operation := func() (interface{}, error) {
		return service()
	}

	bo := backoff.NewExponentialBackOff()
	bo.MaxInterval = 10 * time.Second

	result, err := backoff.Retry(context.TODO(), operation, backoff.WithBackOff(bo))

	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return result, nil
}
