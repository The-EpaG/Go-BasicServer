package e2e_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	test_utils "github.com/The-EpaG/Go-BasicServer/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	t.Run("positive_GET /version", func(t *testing.T) {
		res, err := http.Get(test_utils.Url + "/version")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, res.StatusCode)
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, strings.TrimSpace(string(body)))
	})

	t.Run("positive_GET /health", func(t *testing.T) {
		res, err := http.Get(test_utils.Url + "/health")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, res.StatusCode)
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Empty(t, strings.TrimSpace(string(body)))
	})

}
