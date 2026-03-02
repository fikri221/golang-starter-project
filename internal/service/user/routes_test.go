package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jwt-auth/internal/types"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserServiceHandler(t *testing.T) {
	// TODO: do mock test for user service handler
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserRequest{
			Email:     "",
			Password:  "password",
			FirstName: "John",
			LastName:  "Doe",
		}

		jsonPayload, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonPayload))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		handler.handleRegister(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Should successfully register a user", func(t *testing.T) {
		payload := types.RegisterUserRequest{
			Email:     "fikri@mail.com",
			Password:  "password",
			FirstName: "John",
			LastName:  "Doe",
		}

		jsonPayload, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonPayload))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		handler.handleRegister(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
