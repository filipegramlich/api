package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Max":    20,
			"Filipe": 10,
		},
	}

	server := &PlayerServer{&store}
	t.Run("returns Max's score", func(t *testing.T) {
		request := newGetScoreRequest("Max")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Filipe's score", func(t *testing.T) {
		request := newGetScoreRequest("Filipe")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
