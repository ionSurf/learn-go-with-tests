package contexts

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	// slow response fetching simulation - letter by letter from response
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy got cancelled")
				return
			default:
				time.Sleep(100 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Errorf("store was not told to be cancelled")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Errorf("store was told to be cancelled")
// 	}
// }

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "Hello, World"
		store := NewSpyStore(data, t)
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}

		// store.assertWasNotCancelled()
	})
	t.Run("cancel the request", func(t *testing.T) {
		data := "Hello, World"
		store := NewSpyStore(data, t)
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := NewSpyResponseWriter()

		srv.ServeHTTP(response, request)

		// store.assertWasCancelled()
		if response.written {
			t.Errorf("Response should not have been written")
		}
	})
}

func NewSpyStore(response string, t *testing.T) *SpyStore {
	return &SpyStore{response: response, t: t}
}

func NewSpyResponseWriter() *SpyResponseWriter {
	return &SpyResponseWriter{written: false}
}
