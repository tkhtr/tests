package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			userResult, err := ioutil.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)

			var user User
			err = json.Unmarshal(userResult, &user)
			require.NoError(t, err)

			assert.Equal(t, tt.want.user, user)
		})
	}
}
