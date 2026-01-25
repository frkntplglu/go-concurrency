package healthchecker

import (
	"context"
	"fmt"
	"sync"
)

type client interface {
	Get(ctx context.Context, url string, headers map[string]string) (int, error)
}

type HealthChecker struct {
	urls   []string
	client client
}

func New(urls []string, client client) *HealthChecker {
	return &HealthChecker{
		urls:   urls,
		client: client,
	}
}

func (h *HealthChecker) CheckSequentially() {
	for _, url := range h.urls {
		statusCode := h.check(context.Background(), url)
		fmt.Printf("URL : %s, Status : %d \n", url, statusCode)
	}
}

func (h *HealthChecker) CheckConcurrently(ctx context.Context) {
	var wg sync.WaitGroup

	for _, url := range h.urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			statusCode := h.check(ctx, url)
			fmt.Printf("URL : %s, Status : %d \n", url, statusCode)
		}(url)
	}

	wg.Wait()

}

func (h *HealthChecker) check(ctx context.Context, url string) int {
	statusCode, _ := h.client.Get(ctx, url, nil)

	return statusCode
}
