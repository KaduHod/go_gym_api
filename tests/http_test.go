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
    t.Run("Unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/muscles/groups", nil)
		req.Header.Add("Authorization", "Bearer wrong_token")
		resp, _ := client.Do(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
    })
    t.Run("Unauthorized without token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/muscles/groups", nil)
		resp, _ := client.Do(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
    })
    t.Run("Unauthorized invalid header", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/muscles/groups", nil)
		req.Header.Add("Authorization", "wrong_token")
		resp, _ := client.Do(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
    })
    t.Run("Unauthorized header without bearer string", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/muscles/groups", nil)
		req.Header.Add("Authorization", "IkaCFL4eYMjysMQtW3dcIQ==:S2FkdUhvZA==")
		resp, _ := client.Do(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
    })
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


	t.Run("ListExercises - Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/exercises", nil)
		req.Header.Add("Authorization", authToken)
		resp, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("ListExercisesAmm - Success ", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/exercises/84", nil)
		req.Header.Add("Authorization", authToken)
		resp, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
