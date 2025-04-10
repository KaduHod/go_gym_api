package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:3005/api/v1" // Change to your dev server URL
const authToken = "Bearer IkaCFL4eYMjysMQtW3dcIQ==:S2FkdUhvZA=="

func TestMusculoSkeletalRoutes(t *testing.T) {
	client := &http.Client{}

	t.Run("ListMuscleGroups - Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/muscles/groups", nil)
		req.Header.Add("Authorization", authToken)
		resp, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("ListMusclePortions - Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/muscles/portions", nil)
		req.Header.Add("Authorization", authToken)
		resp, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("ListJoints - Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/joints", nil)
		req.Header.Add("Authorization", authToken)
		resp, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("ListMovements - Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/movements", nil)
		req.Header.Add("Authorization", authToken)
		resp, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("ListAmm - Success with filters", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/muscles/movement-map?muscle_group=test", nil)
		req.Header.Add("Authorization", authToken)
		resp, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
