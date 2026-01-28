package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserViewHandler(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		user        User
	}

	tests := []struct {
		name    string
		request string
		users   map[string]User
		want    want
	}{
		{
			name:    "Test #1",
			request: "/users?user_id=id1",
			users: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Maral",
					LastName:  "Bizhanov",
				},
			},
			want: want{
				contentType: "application/json",
				statusCode:  200,
				user: User{
					ID:        "id1",
					FirstName: "Maral",
					LastName:  "Bizhanov",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest("POST", tt.request, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(UserViewHandler(tt.users))
			h(w, request)

			result := w.Result()

			assert := assert.New(t)
			require := require.New(t)

			assert.Equal(tt.want.statusCode, result.StatusCode)
			assert.Equal(tt.want.contentType, result.Header.Get("Content-Type"))

			userResult, err := io.ReadAll(result.Body)
			require.NoError(err)
			err = result.Body.Close()
			require.NoError(err)

			var user User
			err = json.Unmarshal(userResult, &user)
			require.NoError(err)

			assert.Equal(tt.want.user, user)
		})
	}
}

func TestUserCreateHandler(t *testing.T) {
	type want struct {
		contentType string
		statusCode int
		user User
	}
	
	tests := []struct{
		name string
		method string
		request string
		body string
		users map[string]User
		want want
	}{
		{
			name: "Test #1: create user successfully",
			method: http.MethodPost,
			request: "/users",
			body: `{
				"id": "id1",
				"FirstName": "Maral",
				"LastName": "Bizhanov"
			}`,
			users: map[string]User{},
			want: want{
				contentType: "application/json",
				statusCode: http.StatusCreated,
				user: User{
					ID: "id1",
					FirstName: "Maral",
					LastName: "Bizhanov",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.request, strings.NewReader(tt.body))
			request.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			h := UserCreateHandler(tt.users)
			h(w, request)

			assert := assert.New(t)

			result := w.Result()
			defer result.Body.Close()

			assert.Equal(tt.want.contentType, result.Header.Get("Content-type"))
			assert.Equal(tt.want.statusCode, result.StatusCode)

			var user User
			err := json.NewDecoder(result.Body).Decode(&user)
			require.NoError(t, err)

			assert.Equal(tt.want.user, user)
		})
	}
}
