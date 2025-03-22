package e2e_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	_ "embed"

	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models"
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models/dtos"
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/repository"
	"github.com/The-EpaG/Go-BasicServer/test/mocks"
	test_utils "github.com/The-EpaG/Go-BasicServer/test/utils"
	"github.com/stretchr/testify/assert"
)

var userDTO dtos.UserDTO

func TestUser(t *testing.T) {

	t.Run("positive_GET /users", func(t *testing.T) {
		res, err := http.Get(test_utils.Url + "/users")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, res.StatusCode)
		defer res.Body.Close()
		rawBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		var body []map[string]any
		err = json.Unmarshal(rawBody, &body)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, body)
		user := body[0]
		createdAtStr, ok := user["created_at"].(string)
		if !ok {
			t.Fatal("created_at is not a string")
		}
		createdAt, err := time.Parse(time.RFC3339, createdAtStr)
		if err != nil {
			t.Fatal(err)
		}
		updatedAtStr, ok := user["updated_at"].(string)
		if !ok {
			t.Fatal("updated_at is not a string")
		}
		updatedAt, err := time.Parse(time.RFC3339, updatedAtStr)
		if err != nil {
			t.Fatal(err)
		}
		userDTO = dtos.UserDTO{
			Uuid:      user["uuid"].(string),
			Name:      user["name"].(string),
			Username:  user["username"].(string),
			Surname:   user["surname"].(string),
			Email:     user["email"].(string),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		assert.Equal(t, "admin", user["name"])
	})

	t.Run("positive_GET /users/:id", func(t *testing.T) {
		res, err := http.Get(test_utils.Url + "/users/" + userDTO.Uuid)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, res.StatusCode)
		defer res.Body.Close()
		rawBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		var body map[string]any
		err = json.Unmarshal(rawBody, &body)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, body)
		assert.Equal(t, userDTO.Uuid, body["uuid"])
		assert.Equal(t, userDTO.Name, body["name"])
		assert.Equal(t, userDTO.Username, body["username"])
		assert.Equal(t, userDTO.Surname, body["surname"])
		assert.Equal(t, userDTO.Email, body["email"])
		createdAtStr, ok := body["created_at"].(string)
		if !ok {
			t.Fatal("created_at is not a string")
		}
		createdAt, err := time.Parse(time.RFC3339, createdAtStr)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, userDTO.CreatedAt.Format(time.RFC3339), createdAt.Format(time.RFC3339))
		updatedAtStr, ok := body["updated_at"].(string)
		if !ok {
			t.Fatal("updated_at is not a string")
		}
		updatedAt, err := time.Parse(time.RFC3339, updatedAtStr)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, userDTO.UpdatedAt.Format(time.RFC3339), updatedAt.Format(time.RFC3339))
	})

	t.Run("negative_GET /users/:id - ID not found", func(t *testing.T) {
		const userId = "3caadebd-9dcf-4752-bc14-dd720e16556e"
		res, err := http.Get(test_utils.Url + "/users/" + userId)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		defer res.Body.Close()
		rawBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		var body map[string]any
		err = json.Unmarshal(rawBody, &body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, float64(1), body["code"])
	})

	t.Run("negative_GET /users/:id - invalid ID", func(t *testing.T) {
		res, err := http.Get(test_utils.Url + "/users/348236gghbhkwbfhwb")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		defer res.Body.Close()
		rawBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		var body map[string]any
		err = json.Unmarshal(rawBody, &body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, float64(2), body["code"])
	})

	t.Run("positive_GET /users/external", func(t *testing.T) {
		userExternalRes, err := test_utils.NewJsonBuilder().
			SetArray("data", []any{
				map[string]any{
					"id":         7,
					"email":      "michael.lawson@reqres.in",
					"first_name": "Michael",
					"last_name":  "Lawson",
					"avatar":     "https://reqres.in/img/faces/7-image.jpg",
				},
			},
			).
			SetInt("total", 7).
			SetInt("page", 2).
			SetInt("per_page", 6).
			SetInt("total_pages", 2).
			Build()
		if err != nil {
			t.Fatal(err)
		}
		mocks.HttpGetReturnValue = userExternalRes
		mocks.HttpGetErrorValue = nil

		res, err := http.Get(test_utils.Url + "/users/external")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, res.StatusCode)
		defer res.Body.Close()
		rawBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		var body []map[string]any
		err = json.Unmarshal(rawBody, &body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 1, len(body))
	})

	t.Run("negative_GET /users/external - external service error", func(t *testing.T) {
		mocks.HttpGetReturnValue = []byte{}
		mocks.HttpGetErrorValue = errors.New("server is down")

		res, err := http.Get(test_utils.Url + "/users/external")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusServiceUnavailable, res.StatusCode)
		defer res.Body.Close()
		rawBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		var body map[string]any
		err = json.Unmarshal(rawBody, &body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, float64(3), body["code"])
	})

	t.Run("positive_POST /users/", func(t *testing.T) {
		container.Invoke(func(userRepository repository.UsersRepository) {
			createDto := dtos.CreateUserDTO{
				Username: "test",
				Name:     "test",
				Surname:  "test",
				Email:    "test@test.it",
				Password: "abc",
			}
			dto, err := json.Marshal(createDto)
			if err != nil {
				t.Fatal(err)
			}
			res, err := http.Post(test_utils.Url+"/users/", "application/json", bytes.NewBuffer(dto))
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, http.StatusCreated, res.StatusCode)
			defer res.Body.Close()
			rawBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			var body map[string]any
			err = json.Unmarshal(rawBody, &body)
			if err != nil {
				t.Fatal(err)
			}

			assert.NotEmpty(t, body)
			assert.Equal(t, body["email"], "test@test.it")
			assert.Equal(t, body["name"], "test")
			assert.Equal(t, body["surname"], "test")
			assert.Nil(t, body["password"])

			users := userRepository.FindAll()
			assert.Equal(t, 2, len(users))
			uuid := body["uuid"].(string)
			// cleanup
			userRepository.SoftDeleteUser(context.Background(), uuid)
		})
	})

	t.Run("negative_POST /users/ - duplicate email", func(t *testing.T) {
		container.Invoke(func(userRepository repository.UsersRepository) {
			username := "test"
			userRepository.CreateUser(context.Background(), models.User{
				Username:     &username,
				Name:         "test",
				Surname:      "test",
				Email:        "test@test.it",
				PasswordHash: "abc",
			})

			createDto := dtos.CreateUserDTO{
				Username: "test",
				Name:     "test",
				Surname:  "test",
				Email:    "test@test.it",
				Password: "abc",
			}
			dto, err := json.Marshal(createDto)
			if err != nil {
				t.Fatal(err)
			}
			res, err := http.Post(test_utils.Url+"/users/", "application/json", bytes.NewBuffer(dto))
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, http.StatusConflict, res.StatusCode)
			defer res.Body.Close()
			rawBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			var body map[string]any
			err = json.Unmarshal(rawBody, &body)
			if err != nil {
				t.Fatal(err)
			}

			assert.NotEmpty(t, body)
			assert.Equal(t, float64(4), body["code"])

			users := userRepository.FindAll()
			assert.Equal(t, 2, len(users))

			uuid := users[1].Uuid
			// cleanup
			userRepository.SoftDeleteUser(context.Background(), uuid.String())
		})
	})
}
