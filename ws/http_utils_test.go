package ws

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type DummyData struct {
	Message string `json:"message"`
}

var (
	server      *httptest.Server
	testBaseURL string
)

func init() {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/success":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"message":"Hello World!"}`))

		case "/invalid-json":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{invalid json}`))

		case "/error":
			w.WriteHeader(http.StatusInternalServerError)

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	testBaseURL = server.URL
}

func TestGet_Success(t *testing.T) {
	client := server.Client()
	expected := DummyData{Message: "Hello World!"}

	result, err := get[DummyData](client, testBaseURL+"/success")

	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestGet_InvalidJSON(t *testing.T) {
	client := server.Client()

	_, err := get[DummyData](client, testBaseURL+"/invalid-json")

	require.Error(t, err)
}

func TestGet_Non200Status(t *testing.T) {
	client := server.Client()

	_, err := get[DummyData](client, testBaseURL+"/error")

	require.Error(t, err)
}
