package healthchecker

import (
	"context"
	"errors"
	"testing"
)

type MockClient struct {
	MockGet func(ctx context.Context, url string) (int, error)
}

func (m *MockClient) Get(ctx context.Context, url string, headers map[string]string) (int, error) {
	return m.MockGet(ctx, url)
}

func TestCheckConcurrently(t *testing.T) {
	urls := []string{"https://test1.com", "https://test2.com"}

	mock := &MockClient{
		MockGet: func(ctx context.Context, url string) (int, error) {
			if url == "https://test1.com" {
				return 200, nil
			}
			return 404, errors.New("not found")
		},
	}

	checker := New(urls, mock)

	t.Run("should complete without panic", func(t *testing.T) {
		ctx := context.Background()
		checker.CheckConcurrently(ctx)
	})
}
